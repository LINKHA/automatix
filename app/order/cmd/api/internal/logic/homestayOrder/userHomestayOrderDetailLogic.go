package homestayOrder

import (
	"context"

	"github.com/LINKHA/automatix/app/order/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/order/cmd/api/internal/types"
	"github.com/LINKHA/automatix/app/order/cmd/rpc/order"
	"github.com/LINKHA/automatix/app/order/model"
	"github.com/LINKHA/automatix/app/payment/cmd/rpc/payment"
	"github.com/LINKHA/automatix/common/ctxdata"
	"github.com/LINKHA/automatix/common/tool"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserHomestayOrderDetailLogic {
	return UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req types.UserHomestayOrderDetailReq) (*types.UserHomestayOrderDetailResp, error) {

	userId := ctxdata.GetUidFromCtx(l.ctx)

	resp, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay order detail fail"), " rpc get HomestayOrderDetail err:%v , sn : %s", err, req.Sn)
	}

	var typesOrderDetail types.UserHomestayOrderDetailResp
	if resp.HomestayOrder != nil && resp.HomestayOrder.UserId == userId {

		copier.Copy(&typesOrderDetail, resp.HomestayOrder)

		//重置价格.
		typesOrderDetail.OrderTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.OrderTotalPrice)
		typesOrderDetail.FoodTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodTotalPrice)
		typesOrderDetail.HomestayTotalPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayTotalPrice)
		typesOrderDetail.HomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.HomestayPrice)
		typesOrderDetail.FoodPrice = tool.Fen2Yuan(resp.HomestayOrder.FoodPrice)
		typesOrderDetail.MarketHomestayPrice = tool.Fen2Yuan(resp.HomestayOrder.MarketHomestayPrice)

		//支付信息.
		if typesOrderDetail.TradeState != model.HomestayOrderTradeStateCancel && typesOrderDetail.TradeState != model.HomestayOrderTradeStateWaitPay {
			paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentSuccessRefundByOrderSn(l.ctx, &payment.GetPaymentSuccessRefundByOrderSnReq{
				OrderSn: resp.HomestayOrder.Sn,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Failed to get order payment information err : %v , orderSn:%s", err, resp.HomestayOrder.Sn)
			}

			if paymentResp.PaymentDetail != nil {
				typesOrderDetail.PayTime = paymentResp.PaymentDetail.PayTime
				typesOrderDetail.PayType = paymentResp.PaymentDetail.PayMode
			}
		}

		return &typesOrderDetail, nil
	}

	return nil, nil

}
