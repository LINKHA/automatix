package logic

import (
	"context"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServerListLogic {
	return &GetServerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServerListLogic) GetServerList(in *pb.GetServerListReq) (*pb.GetServerListResp, error) {
	whereBuilder := l.svcCtx.ServerModel.SelectBuilder().Where(
		"server_type = ? and switch = ?",
		0, 1,
	)

	serverList, err := l.svcCtx.ServerModel.FindAll(l.ctx, whereBuilder, "id desc")
	if err != nil {
		return &pb.GetServerListResp{ReturnCode: int64(xerr.SERVER_COMMON_ERROR)}, nil
	}

	var serverListResp []*pb.ServerInfo
	if len(serverList) > 0 {
		for _, serverInfo := range serverList {
			var pbServerInfo pb.ServerInfo
			_ = copier.Copy(&pbServerInfo, serverInfo)
			serverListResp = append(serverListResp, &pbServerInfo)
		}
	}

	return &pb.GetServerListResp{
		ReturnCode: int64(xerr.OK),
		ServerList: serverListResp,
	}, nil
}
