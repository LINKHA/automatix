package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleLogic) CreateRole(in *pb.CreateRoleReq) (*pb.CreateRoleResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateRoleResp{}, nil
}
