package logic

import (
	"context"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/common/servercode"
	"automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnterServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnterServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterServerLogic {
	return &EnterServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EnterServerLogic) EnterServer(in *pb.EnterServerReq) (*pb.EnterServerResp, error) {

	_, serverId := servercode.UseServerCode(l.svcCtx.Redis, in.ServerCode)
	if serverId == "" {
		return &pb.EnterServerResp{
			ReturnCode: int64(xerr.SERVER_CODE_INVALID),
			Host:       "",
			Port:       "",
		}, nil
	}

	return &pb.EnterServerResp{
		ReturnCode: int64(xerr.OK),
		Host:       "127.0.0.1",
		Port:       "1111",
	}, nil
}
