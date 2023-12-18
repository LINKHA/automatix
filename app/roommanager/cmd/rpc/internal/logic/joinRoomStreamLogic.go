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
	msgBox chan pb.JoinRoomStreamResp
}

func NewJoinRoomStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomStreamLogic {
	return &JoinRoomStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		msgBox: make(chan pb.JoinRoomStreamResp, 1000),
	}
}

func (l *JoinRoomStreamLogic) JoinRoomStream(stream pb.Roommanager_JoinRoomStreamServer) error {

	select {
	case data := <-l.msgBox:
		stream.Send(&data)
		// stream.Send(&pb.JoinRoomStreamResp{
		// 	ReturnCode: 1,
		// })
	case <-l.ctx.Done():
		return nil
	}
	return nil
}

func (l *JoinRoomStreamLogic) SendJoinRoomStream(joinRoomStreamResp pb.JoinRoomStreamResp) {
	l.msgBox <- joinRoomStreamResp
}
