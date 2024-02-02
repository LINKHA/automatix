package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoomLogic) CreateRoom(stream pb.Roommanager_CreateRoomServer) error {
	go func() {
		for {
			select {
			case <-l.ctx.Done():
				return
			default:
				msg, err := stream.Recv()
				fmt.Println(err)
				l.handlerFunc(stream, msg)

			}
		}
	}()

	select {
	case <-l.ctx.Done():
		return nil
	}
}

func (l *CreateRoomLogic) handlerFunc(stream pb.Roommanager_CreateRoomServer, req *pb.CreateRoomReq) {

}
