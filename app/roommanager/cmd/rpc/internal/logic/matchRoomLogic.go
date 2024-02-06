package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MatchRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMatchRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MatchRoomLogic {
	return &MatchRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MatchRoomLogic) MatchRoom(stream pb.Roommanager_MatchRoomServer) error {
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

func (l *MatchRoomLogic) handlerFunc(stream pb.Roommanager_MatchRoomServer, req *pb.MatchRoomReq) {
}
