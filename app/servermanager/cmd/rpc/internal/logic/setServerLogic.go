package logic

import (
	"context"
	"fmt"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/app/servermanager/model"
	"automatix/common/xerr"

	"github.com/pkg/errors"
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
	server := model.Server{
		ServerId:   in.ServerId,
		Name:       in.Name,
		ServerType: in.ServerType,
		Switch:     in.Switch,
		StartTime:  in.StartTime,
		// Area: in.Area,
		Tags:          in.Tags,
		MaxOnline:     in.MaxOnline,
		MaxQueue:      in.MaxQueue,
		MaxSign:       in.MaxSign,
		TemplateValue: in.TemplateValue,
	}
	fmt.Print("1---------    :   ")
	fmt.Println(server)

	fmt.Print("2---------    :   ")
	fmt.Println(in.Switch)
	err := l.svcCtx.ServerModel.UpdateByServerId(l.ctx, &server)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), " SetServer db err , serverId : %s ", in.ServerId)
	}

	return &pb.SetServerResp{ReturnCode: 0}, nil
}
