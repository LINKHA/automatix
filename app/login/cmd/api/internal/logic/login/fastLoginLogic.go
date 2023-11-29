package login

import (
	"context"

	"automatix/app/login/cmd/api/internal/svc"
	"automatix/app/login/cmd/api/internal/types"

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

	return
}
