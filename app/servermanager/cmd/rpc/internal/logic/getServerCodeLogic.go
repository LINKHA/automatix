package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/servercode"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServerCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerCodeLogic {
	return &GetServerCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServerCodeLogic) GetServerCode(in *pb.GetServerCodeReq) (*pb.GetServerCodeResp, error) {
	serverCode := servercode.FindServerCode(l.svcCtx.Redis, in.UserId, in.ServerId)
	return &pb.GetServerCodeResp{
		ReturnCode: int64(xerr.OK),
		ServerCode: serverCode,
	}, nil
}
