package logic

import (
	"context"
	"strconv"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/app/servermanager/model"
	"github.com/LINKHA/automatix/common/xerr"

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
	var returnCode = xerr.OK

	server := new(model.Server)
	serverId := strconv.FormatInt(int64(l.svcCtx.Snowflake.Generate()), 10)

	server.ServerId = serverId
	server.Name = in.Name
	server.ServerType = in.ServerType
	server.Switch = in.Switch
	server.StartTime = in.StartTime
	// server.Area = in.Area
	server.Tags = in.Tags
	server.MaxOnline = in.MaxOnline
	server.MaxQueue = in.MaxQueue
	server.MaxSign = in.MaxSign
	server.TemplateValue = in.TemplateValue

	_, err := l.svcCtx.ServerModel.Insert(l.ctx, server)
	if err != nil {
		returnCode = xerr.SERVER_COMMON_ERROR
	}

	return &pb.CreateServerResp{
		ReturnCode: int64(returnCode),
		ServerId:   serverId,
	}, nil
}
