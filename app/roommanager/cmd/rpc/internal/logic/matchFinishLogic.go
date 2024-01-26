package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return nil
}
