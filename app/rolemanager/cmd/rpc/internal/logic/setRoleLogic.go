package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleLogic {
	return &SetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoleLogic) SetRole(stream pb.Rolemanager_SetRoleServer) error {
	// todo: add your logic here and delete this line

	return nil
}
