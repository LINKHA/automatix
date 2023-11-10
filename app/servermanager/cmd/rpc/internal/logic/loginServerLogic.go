package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"automatix/app/servermanager/cmd/rpc/internal/svc"
	"automatix/app/servermanager/cmd/rpc/pb"
	"automatix/common/kqueue"

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
	fmt.Print("qwe 1---------    :   ")
	fmt.Println("push login msg")
	var message kqueue.LoginServerMessage
	message.PlayerId = in.PlayerId
	message.ServerId = in.ServerId

	jsonData, _ := json.Marshal(message)

	err := l.svcCtx.KqueueServerQueue.Push(string(jsonData))
	if err != nil {
		fmt.Println(err)
	}

	return &pb.LoginServerResp{}, nil
}
