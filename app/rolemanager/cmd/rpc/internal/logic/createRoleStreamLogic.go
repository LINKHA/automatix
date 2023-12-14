package logic

import (
	"context"
	"fmt"

	"automatix/app/rolemanager/cmd/rpc/internal/svc"
	"automatix/app/rolemanager/cmd/rpc/pb"

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
	fmt.Printf("1------------------   :   %d", 111)
	fmt.Println(stream)
	stream.Send(&pb.CreateRoleResp{
		ReturnCode: 1,
	})
	return nil
}
