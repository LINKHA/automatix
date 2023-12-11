package logic

import (
	"context"

	"automatix/app/playermanager/cmd/rpc/internal/svc"
	"automatix/app/playermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleStreamLogic {
	return &CreateRoleStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleStreamLogic) CreateRoleStream(stream pb.Rolemanager_CreateRoleStreamServer) error {
	// todo: add your logic here and delete this line

	return nil
}
