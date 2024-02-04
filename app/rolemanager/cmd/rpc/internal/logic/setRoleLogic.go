package logic

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/model"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetRoleLogic {
	return &SetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetRoleLogic) SetRole(stream pb.Rolemanager_SetRoleServer) error {
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

func (l *SetRoleLogic) handlerFunc(stream pb.Rolemanager_SetRoleServer, req *pb.SetRoleReq) {
	role := new(model.Role)
	role.RoleId = req.RoleId
	role.HistoryServerIds = "[]"
	role.Tags = "[]"
	role.TemplateValue = req.TemplateValue
	err := l.svcCtx.RoleModel.UpdateByRoleId(l.ctx, role)
	if err != nil {
		fmt.Println(err)

		stream.Send(&pb.SetRoleResp{
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}

	stream.Send(&pb.SetRoleResp{
		ReturnCode: int64(xerr.OK),
	})
}
