package logic

import (
	"context"

	"automatix/app/roommanager/cmd/rpc/internal/svc"
	"automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomStreamLogic {
	return &JoinRoomStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinRoomStreamLogic) JoinRoomStream(stream pb.Roommanager_JoinRoomStreamServer) error {
	// todo: add your logic here and delete this line
	stream.Send(&pb.JoinRoomStreamResp{
		ReturnCode: 1,
	})
	return nil
}
