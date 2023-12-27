package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/types"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/servermanager"
	"github.com/LINKHA/automatix/common/ctxdata"

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
	userId := ctxdata.GetUidFromCtx(l.ctx)
	loginServerResp, err := l.svcCtx.ServerManagerRpc.LoginServer(l.ctx, &servermanager.LoginServerReq{
		UserId:   userId,
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	var resp types.LoginServerResp
	_ = copier.Copy(&resp, loginServerResp)

	return &resp, nil
}
