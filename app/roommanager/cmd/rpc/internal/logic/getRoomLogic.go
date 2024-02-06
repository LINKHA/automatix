package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomLogic {
	return &GetRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoomLogic) GetRoom(stream pb.Roommanager_GetRoomServer) error {
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

func (l *GetRoomLogic) handlerFunc(stream pb.Roommanager_GetRoomServer, req *pb.GetRoomReq) {
}
