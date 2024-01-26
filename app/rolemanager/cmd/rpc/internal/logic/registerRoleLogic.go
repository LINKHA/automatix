package logic

import (
	"context"
	"fmt"

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

	stream.Send(&pb.RegisterRoleResp{
		RoldId:     "123",
		ReturnCode: 2,
	})

	//stream reader
	go func() {
		for {
			select {
			case <-l.ctx.Done():
				return
			default:
				gateMsg, err := stream.Recv()
				fmt.Println(gateMsg)
				fmt.Println(err)
			}
		}
	}()

	select {}
	return nil
}
