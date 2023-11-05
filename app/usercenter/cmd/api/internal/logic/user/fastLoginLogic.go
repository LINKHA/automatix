package user

import (
	"context"
	"fmt"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FastLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFastLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FastLoginLogic {
	return &FastLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FastLoginLogic) FastLogin(req *types.FastLoginReq) (resp *types.FastLoginResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("fast login+++++++++++++++++++++++")
	// loginResp, err := l.svcCtx.UsercenterRpc.Login(l.ctx, &usercenter.LoginReq{
	// 	AuthType: model.UserAuthTypeSystem,
	// 	AuthKey:  req.Mobile,
	// 	Password: req.Password,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	// var resp types.LoginResp
	// _ = copier.Copy(&resp, loginResp)

	//return &resp, nil
	return
}
