package logic

import (
	"context"

	"looklook/app/servermanager/cmd/rpc/internal/svc"
	"looklook/app/servermanager/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/pkg/errors"
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
	server, err := l.svcCtx.ServerModel.FindOneByServerId(l.ctx, in.ServerId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Get db server Insert err:%v,server:%+v", err, server)
	}
	return &pb.GetServerResp{
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
