package login

import (
	"context"

	"automatix/app/login/cmd/api/internal/svc"
	"automatix/app/login/cmd/api/internal/types"
	"automatix/app/servermanager/cmd/rpc/servermanager"
	"automatix/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerCodeLogic {
	return &GetServerCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerCodeLogic) GetServerCode(req *types.GetServerCodeReq) (*types.GetServerCodeResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	getServerCodeResp, err := l.svcCtx.ServerManagerRpc.GetServerCode(l.ctx, &servermanager.GetServerCodeReq{
		UserId:   userId,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	var resp types.GetServerCodeResp
	_ = copier.Copy(&resp, getServerCodeResp)

	return &resp, nil
}
