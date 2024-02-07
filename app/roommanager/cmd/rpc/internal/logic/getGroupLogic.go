package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupLogic {
	return &GetGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupLogic) GetGroup(stream pb.Roommanager_GetGroupServer) error {
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

func (l *GetGroupLogic) handlerFunc(stream pb.Roommanager_GetGroupServer, req *pb.GetGroupReq) {
	groupKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_GROUP, req.GroupId)
	groupStr, err := l.svcCtx.Redis.Get(groupKey)
	if err != nil {
		stream.Send(&pb.GetGroupResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
		return
	}

	m := &svc.Group{}
	json.Unmarshal([]byte(groupStr), m)

	stream.Send(&pb.GetGroupResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
		GroupId:    m.GroupID,
		GroupName:  m.GroupName,
		MaxPlayer:  int64(m.MaxPlayer),
		RoleIds:    m.Roles,
	})
}
