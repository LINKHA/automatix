package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
