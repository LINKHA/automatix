package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/api/internal/types"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/servermanager"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerListLogic {
	return &GetServerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerListLogic) GetServerList(req *types.GetServerListReq) (*types.GetServerListResp, error) {
	getServerListResp, err := l.svcCtx.ServerManagerRpc.GetServerList(l.ctx, &servermanager.GetServerListReq{})

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var resp types.GetServerListResp

	_ = copier.Copy(&resp, getServerListResp)

	return &resp, nil
}
