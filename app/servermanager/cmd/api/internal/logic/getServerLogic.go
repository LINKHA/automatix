package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/types"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/servermanager"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerLogic {
	return &GetServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerLogic) GetServer(req *types.GetServerReq) (*types.GetServerResp, error) {
	getServerResp, err := l.svcCtx.ServerManagerRpc.GetServer(l.ctx, &servermanager.GetServerReq{
		ServerId: req.ServerId,
	})
	if err != nil {
		return nil, err
	}

	var resp types.GetServerResp

	_ = copier.Copy(&resp, getServerResp)

	return &resp, nil
}
