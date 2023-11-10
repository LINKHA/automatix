package logic

import (
	"context"

	"automatix/app/servermanager/cmd/api/internal/svc"
	"automatix/app/servermanager/cmd/api/internal/types"
	"automatix/app/servermanager/cmd/rpc/servermanager"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginServerLogic {
	return &LoginServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginServerLogic) LoginServer(req *types.LoginServerReq) (*types.LoginServerResp, error) {

	loginServerResp, err := l.svcCtx.ServerManagerRpc.LoginServer(l.ctx, &servermanager.LoginServerReq{
		PlayerId: req.PlayerId,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	var resp types.LoginServerResp
	_ = copier.Copy(&resp, loginServerResp)

	return &resp, nil
}
