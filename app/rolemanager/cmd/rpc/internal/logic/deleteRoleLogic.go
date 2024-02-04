package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleLogic {
	return &DeleteRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleLogic) DeleteRole(stream pb.Rolemanager_DeleteRoleServer) error {
	go func() {
		for {
			select {
			case <-l.ctx.Done():
				return
			default:
				msg, err := stream.Recv()
				fmt.Println(err)
				l.handlerFunc(stream, msg)

			}
		}
	}()

	select {
	case <-l.ctx.Done():
		return nil
	}
}

func (l *DeleteRoleLogic) handlerFunc(stream pb.Rolemanager_DeleteRoleServer, req *pb.DeleteRoleReq) {
	err := l.svcCtx.RoleModel.DeleteByRoleId(l.ctx, req.RoleId)
	fmt.Println("1-----------------", err)

	if err != nil {
		fmt.Println(err)

		stream.Send(&pb.DeleteRoleResp{
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}

	stream.Send(&pb.DeleteRoleResp{
		ReturnCode: int64(xerr.OK),
	})
}
