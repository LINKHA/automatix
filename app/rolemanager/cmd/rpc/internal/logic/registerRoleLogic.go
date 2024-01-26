package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterRoleLogic {
	return &RegisterRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterRoleLogic) RegisterRole(stream pb.Rolemanager_RegisterRoleServer) error {
	// todo: add your logic here and delete this line

	return nil
}
