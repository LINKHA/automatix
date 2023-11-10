package logic

import (
	"context"
	"encoding/json"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/common/kqueue"
	"automatix/common/xerr"

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
	var returnCode = xerr.OK
	var message kqueue.LoginServerMessage
	message.PlayerId = in.PlayerId
	message.ServerId = in.ServerId

	jsonData, _ := json.Marshal(message)

	err := l.svcCtx.KqueueServerQueue.Push(string(jsonData))
	if err != nil {
		returnCode = xerr.SERVER_COMMON_ERROR
	}

	return &pb.LoginServerResp{
		ReturnCode: int64(returnCode),
	}, nil
}
