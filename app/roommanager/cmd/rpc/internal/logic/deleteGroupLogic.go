package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(stream pb.Roommanager_DeleteGroupServer) error {
	// todo: add your logic here and delete this line

	return nil
}
