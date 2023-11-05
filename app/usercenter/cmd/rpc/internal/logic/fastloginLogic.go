package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FastloginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFastloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastloginLogic {
	return &FastloginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FastloginLogic) Fastlogin(in *pb.FastLoginReq) (*pb.FastLoginResp, error) {
	// todo: add your logic here and delete this line

	return &pb.FastLoginResp{}, nil
}
