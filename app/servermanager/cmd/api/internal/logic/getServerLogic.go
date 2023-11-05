package logic

import (
	"context"

	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"

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

func (l *GetServerLogic) GetServer(req *types.GetServerReq) (resp *types.GetServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
