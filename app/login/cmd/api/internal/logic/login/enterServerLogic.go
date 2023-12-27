package login

import (
	"context"

	"github.com/LINKHA/automatix/app/login/cmd/api/internal/svc"
	"github.com/LINKHA/automatix/app/login/cmd/api/internal/types"
	"github.com/LINKHA/automatix/app/servermanager/cmd/rpc/servermanager"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type EnterServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnterServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnterServerLogic {
	return &EnterServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnterServerLogic) EnterServer(req *types.EnterServerReq) (*types.EnterServerResp, error) {
	enterServerCodeResp, err := l.svcCtx.ServerManagerRpc.EnterServer(l.ctx, &servermanager.EnterServerReq{
		ServerCode: req.ServerCode,
	})
	if err != nil {
		return nil, err
	}

	var resp types.EnterServerResp
	_ = copier.Copy(&resp, enterServerCodeResp)

	return &resp, nil
}
