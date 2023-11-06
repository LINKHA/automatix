package logic

import (
	"context"
	"fmt"

	"looklook/app/servermanager/cmd/rpc/internal/svc"
	"looklook/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServerLogic {
	return &CreateServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateServerLogic) CreateServer(in *pb.CreateServerReq) (*pb.CreateServerResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("CreateServer++++++++++++++++")
	return &pb.CreateServerResp{}, nil
}
