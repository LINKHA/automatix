package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomLogic {
	return &JoinRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinRoomLogic) JoinRoom(in *pb.JoinRoomReq) (*pb.JoinRoomResp, error) {
	// todo: add your logic here and delete this line
	value, _ := l.svcCtx.StreamManager.Get("JoinRoomStream")
	joinRoomStreamLogic := value.(JoinRoomStreamLogic)
	joinRoomStreamLogic.SendJoinRoomStream(&pb.JoinRoomStreamResp{
		ReturnCode: 1,
	})
	return &pb.JoinRoomResp{}, nil
}
