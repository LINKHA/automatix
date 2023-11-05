package logic

import (
	"context"
	"fmt"

	"looklook/app/servermanager/cmd/api/internal/svc"
	"looklook/app/servermanager/cmd/api/internal/types"

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

func (l *CreateServerLogic) CreateServer(req *types.CreateServerReq) (resp *types.CreateServerResp, err error) {
	fmt.Println("Create server++++++++++++++++")
	// todo: add your logic here and delete this line

	return
}
