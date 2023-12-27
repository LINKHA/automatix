package logic

import (
	"context"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

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
	var returnCode = xerr.OK

	server, err := l.svcCtx.ServerModel.FindOneByServerId(l.ctx, in.ServerId)
	if err != nil {
		returnCode = xerr.SERVER_COMMON_ERROR
	}

	return &pb.GetServerResp{
		ReturnCode: int64(returnCode),
		ServerId:   server.ServerId,
		Name:       server.Name,
		ServerType: server.ServerType,
		Switch:     server.Switch,
		StartTime:  server.StartTime,
		// Area:          server.Area,
		Tags:          server.Tags,
		MaxOnline:     server.MaxOnline,
		MaxQueue:      server.MaxQueue,
		MaxSign:       server.MaxSign,
		TemplateValue: server.TemplateValue,
	}, nil
}
