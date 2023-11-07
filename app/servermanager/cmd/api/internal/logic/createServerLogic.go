package logic

import (
	"context"
	"fmt"

	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"
	"looklook/app/servermanager/cmd/rpc/servermanager"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServerLogic {
	return &CreateServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateServerLogic) CreateServer(req *types.CreateServerReq) (*types.CreateServerResp, error) {
	fmt.Println("1 CreateServer++++++++++++++++")

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
	}

	// 生成唯一ID
	id := node.Generate()
	fmt.Print("ggg 1------------------   :   ")
	fmt.Println(id)

	createServerResp, err := l.svcCtx.ServerManagerRpc.CreateServer(l.ctx, &servermanager.CreateServerReq{
		Name:          req.Name,
		ServerType:    req.ServerType,
		Switch:        req.Switch,
		StartTime:     req.StartTime,
		Area:          req.Area,
		Tags:          req.Tags,
		MaxOnline:     req.MaxOnline,
		MaxQueue:      req.MaxOnline,
		MaxSign:       req.MaxSign,
		TemplateValue: req.TemplateValue,
	})
	if err != nil {
		return nil, err
	}

	var resp types.CreateServerResp
	_ = copier.Copy(&resp, createServerResp)

	return &resp, nil
}
