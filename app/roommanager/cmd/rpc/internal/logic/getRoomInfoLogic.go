package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoomInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomInfoLogic {
	return &GetRoomInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoomInfoLogic) GetRoomInfo(stream pb.Roommanager_GetRoomInfoServer) error {
	// todo: add your logic here and delete this line

	return nil
}
