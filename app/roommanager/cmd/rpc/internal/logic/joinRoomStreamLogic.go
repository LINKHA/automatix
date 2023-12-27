package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinRoomStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	msgBox chan *pb.JoinRoomStreamResp
}

func NewJoinRoomStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinRoomStreamLogic {
	joinRoomStreamLogic := &JoinRoomStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		msgBox: make(chan *pb.JoinRoomStreamResp, 1000),
	}

	svcCtx.StreamManager.Set("JoinRoomStream", joinRoomStreamLogic)

	return joinRoomStreamLogic
}

func (l *JoinRoomStreamLogic) JoinRoomStream(stream pb.Roommanager_JoinRoomStreamServer) error {
	for {
		select {
		case data := <-l.msgBox:
			stream.Send(data)
		case <-l.ctx.Done():
			return nil
		}
	}
}

func (l *JoinRoomStreamLogic) SendJoinRoomStream(joinRoomStreamResp *pb.JoinRoomStreamResp) {
	l.msgBox <- joinRoomStreamResp
}
