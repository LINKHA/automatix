package login

import (
	"context"

	"automatix/app/login/cmd/api/internal/svc"
	"automatix/app/login/cmd/api/internal/types"
	"automatix/app/usercenter/cmd/rpc/usercenter"
	"automatix/app/usercenter/model"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FastLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFastLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastLoginLogic {
	return &FastLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FastLoginLogic) FastLogin(req *types.FastLoginReq) (*types.FastLoginResp, error) {
	fastLoginResp, err := l.svcCtx.UsercenterRpc.Fastlogin(l.ctx, &usercenter.FastLoginReq{
		AuthType: model.UserAuthTypeSystem,
		AuthKey:  req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var resp types.FastLoginResp
	_ = copier.Copy(&resp, fastLoginResp)

	return &resp, nil
}
