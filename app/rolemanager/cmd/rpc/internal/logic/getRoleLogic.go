package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(stream pb.Rolemanager_GetRoleServer) error {
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

func (l *GetRoleLogic) handlerFunc(stream pb.Rolemanager_GetRoleServer, req *pb.GetRoleReq) {
	role, err := l.svcCtx.RoleModel.FindOneByRoleId(l.ctx, req.RoleId)
	if err != nil {
		fmt.Println(err)

		stream.Send(&pb.GetRoleResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
		return
	}

	stream.Send(&pb.GetRoleResp{
		Header:           req.Header,
		ReturnCode:       int64(xerr.OK),
		RoleId:           role.RoleId,
		BornServerId:     role.BornServerId,
		CurServerId:      role.CurServerId,
		HistoryServerIds: strings.Split(role.HistoryServerIds, ","),
		Tags:             strings.Split(role.Tags, ","),
		TemplateValue:    role.TemplateValue,
	})
}
