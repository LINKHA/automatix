package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MatchFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMatchFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MatchFinishLogic {
	return &MatchFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MatchFinishLogic) MatchFinish(stream pb.Roommanager_MatchFinishServer) error {
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

func (l *MatchFinishLogic) handlerFunc(stream pb.Roommanager_MatchFinishServer, req *pb.MatchFinishReq) {
}
