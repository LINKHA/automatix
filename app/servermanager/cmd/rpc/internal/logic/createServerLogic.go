package logic

import (
	"context"
	"strconv"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/app/servermanager/model"
	"automatix/common/xerr"

	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Create db server Insert err:%v,server:%+v", err, server)
		// return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Create db server Insert err:%v,server:%+v", err, server)
	}

	return &pb.CreateServerResp{
		ReturnCode: 0,
		ServerId:   serverId,
	}, nil
}
