package logic

import (
	"context"

	"looklook/app/servermanager/cmd/rpc/internal/svc"
	"looklook/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerLogic {
	return &GetServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServerLogic) GetServer(in *pb.GetServerReq) (*pb.GetServerResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetServerResp{}, nil
}
