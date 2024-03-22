package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/friend/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/friend/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(stream pb.Rolemanager_AddFriendServer) error {
	// todo: add your logic here and delete this line

	return nil
}
