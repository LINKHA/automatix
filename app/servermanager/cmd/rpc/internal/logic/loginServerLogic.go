package logic

import (
	"context"
	"encoding/json"

	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/flowlimit"
	"github.com/LINKHA/automatix/common/kqueue"
	"github.com/LINKHA/automatix/common/servercode"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginServerLogic {

	return &LoginServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginServerLogic) LoginServer(in *pb.LoginServerReq) (*pb.LoginServerResp, error) {
	userId := in.UserId
	serverId := in.ServerId

	server, err := l.svcCtx.ServerModel.FindOneByServerId(l.ctx, serverId)

	if err != nil {
		return &pb.LoginServerResp{ReturnCode: int64(xerr.SERVER_COMMON_ERROR)}, nil
	}

	//如果已经存在ServerCode
	findCode := servercode.FindServerCode(l.svcCtx.Redis, userId, serverId)
	if findCode != "" {
		return &pb.LoginServerResp{
			ReturnCode: int64(xerr.OK),
			ServerCode: findCode,
		}, nil
	}

	//If the number of people in line reaches the maximum(如果排队人数到达上限)
	if int(server.MaxQueue) < flowlimit.SlidingWindowCount(l.svcCtx.Redis, serverId) {
		return &pb.LoginServerResp{ReturnCode: int64(xerr.SERVER_MANAGER_LOGIN_SERVER_QUEUE_MAX)}, nil
	}

	//Obtain the traffic limiting configuration(获取限流配置)
	//Note: 这里还缺一个如果kafka中有数据，也需要进入排队
	if flowlimit.SlidingWindow(l.svcCtx.Redis, serverId, int64(l.svcCtx.Config.SlidingWindow.Rate), int64(l.svcCtx.Config.SlidingWindow.WindowSize)) {
		//登录
		serverCode := servercode.GenServerCode(l.svcCtx.Redis, userId, serverId)
		return &pb.LoginServerResp{
			ReturnCode: int64(xerr.OK),
			ServerCode: serverCode,
		}, nil

	} else {
		var message kqueue.LoginServerMessage
		message.UserId = userId
		message.ServerId = serverId
		jsonData, _ := json.Marshal(message)
		err = l.svcCtx.KqueueServerQueue.Push(string(jsonData))

		if err != nil {
			return &pb.LoginServerResp{ReturnCode: int64(xerr.SERVER_COMMON_ERROR)}, nil
		}
		return &pb.LoginServerResp{ReturnCode: int64(xerr.SERVER_MANAGER_LOGIN_SERVER_QUEUE_ENTER)}, nil
	}
}
