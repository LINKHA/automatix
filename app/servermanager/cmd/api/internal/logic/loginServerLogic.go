package logic

import (
	"context"
	"fmt"

	"automatix/app/servermanager/cmd/api/internal/svc"
	"automatix/app/servermanager/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginServerLogic {
	return &LoginServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginServerLogic) LoginServer(req *types.LoginServerReq) (resp *types.LoginServerResp, err error) {
	// var playerId = req.PlayerId
	var serverId = req.ServerId

	errx := l.svcCtx.Kqueue.Push(serverId)
	fmt.Print("0------------------   :   ")
	if errx != nil {
		fmt.Print("1------------------   :   ")
		fmt.Println(errx)
	}
	return
}
