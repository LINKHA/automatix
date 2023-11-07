package logic

import (
	"context"

	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"
	"looklook/app/servermanager/cmd/rpc/servermanager"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetServerLogic {
	return &SetServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetServerLogic) SetServer(req *types.SetServerReq) (*types.SetServerResp, error) {
	setServerResp, err := l.svcCtx.ServerManagerRpc.SetServer(l.ctx, &servermanager.SetServerReq{
		ServerId:   req.ServerId,
		Name:       req.Name,
		ServerType: req.ServerType,
		Switch:     req.Switch,
		// Area:          req.Area,
		Tags:          req.Tags,
		MaxOnline:     req.MaxOnline,
		MaxQueue:      req.MaxOnline,
		MaxSign:       req.MaxSign,
		TemplateValue: req.TemplateValue,
	})
	if err != nil {
		return nil, err
	}

	var resp types.SetServerResp
	_ = copier.Copy(&resp, setServerResp)

	return &resp, nil
}
