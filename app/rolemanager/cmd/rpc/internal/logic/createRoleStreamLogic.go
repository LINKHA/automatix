package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"

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
	fmt.Println("CreateRoleStream Start..")

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

	select {
	case <-l.ctx.Done():
		return nil
	}
}

func (l *CreateRoleStreamLogic) handlerFunc(stream pb.Rolemanager_CreateRoleStreamServer, req *pb.CreateRoleReq) {
	// req.AccountId
	stream.Send(&pb.CreateRoleResp{
		ReturnCode: 2,
	})
}
