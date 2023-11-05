package logic

import (
	"context"

	"looklook/app/servermanager/cmd/rpc/internal/svc"
	"looklook/app/servermanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetServerLogic {
	return &SetServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetServerLogic) SetServer(in *pb.SetServerReq) (*pb.SetServerResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SetServerResp{}, nil
}
