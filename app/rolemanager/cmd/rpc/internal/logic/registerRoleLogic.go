package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/model"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterRoleLogic {
	return &RegisterRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterRoleLogic) RegisterRole(stream pb.Rolemanager_RegisterRoleServer) error {
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

func (l *RegisterRoleLogic) handlerFunc(stream pb.Rolemanager_RegisterRoleServer, req *pb.RegisterRoleReq) {
	role := new(model.Role)
	roleId := strconv.FormatInt(int64(l.svcCtx.Snowflake.Generate()), 10)

	role.RoleId = roleId
	role.BornServerId = req.ServerId
	role.CurServerId = req.ServerId
	role.HistoryServerIds = "[]"
	role.Tags = "[]"
	role.TemplateValue = req.TemplateValue

	_, err := l.svcCtx.RoleModel.Insert(l.ctx, role)

	if err != nil {
		fmt.Println(err)

		stream.Send(&pb.RegisterRoleResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}

	stream.Send(&pb.RegisterRoleResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
		RoleId:     roleId,
	})
}
