package logic

import (
	"context"

	"automatix/app/roommanager/cmd/rpc/internal/svc"
	"automatix/app/roommanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoomLogic {
	return &CreateRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoomLogic) CreateRoom(in *pb.CreateRoomReq) (*pb.CreateRoomResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CreateRoomResp{}, nil
}
