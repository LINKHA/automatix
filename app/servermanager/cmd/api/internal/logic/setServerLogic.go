package logic

import (
	"context"

	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetServerLogic {
	return &SetServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetServerLogic) SetServer(req *types.SetServerReq) (resp *types.SetServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
