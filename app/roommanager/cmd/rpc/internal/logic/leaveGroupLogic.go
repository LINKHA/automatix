package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/tool"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type LeaveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLeaveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveGroupLogic {
	return &LeaveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LeaveGroupLogic) LeaveGroup(stream pb.Roommanager_LeaveGroupServer) error {
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

func (l *LeaveGroupLogic) handlerFunc(stream pb.Roommanager_LeaveGroupServer, req *pb.LeaveGroupReq) {
	err := l.svcCtx.Redis.Pipelined(
		func(pipe redis.Pipeliner) error {
			roleKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_ROLE, req.RoleId)
			groupKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_GROUP, req.GroupId)
			groupStr, err := l.svcCtx.Redis.Get(groupKey)
			if err != nil {
				return err
			}

			m := &svc.Group{}
			json.Unmarshal([]byte(groupStr), m)

			m.Roles = tool.RemoveItemSlice(m.Roles, req.RoleId)

			roleJSON, err2 := json.Marshal(m)
			if err2 != nil {
				return err2
			}
			pipe.SetEX(context.Background(), groupKey, string(roleJSON), 86400*time.Second)

			pipe.HMSet(context.Background(), roleKey, map[string]string{
				"RoleId":  req.RoleId,
				"GroupId": "",
			})

			pipe.Expire(context.Background(), roleKey, 86400*time.Second)
			return nil
		},
	)

	if err != nil {
		stream.Send(&pb.LeaveGroupResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}
	stream.Send(&pb.LeaveGroupResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
	})
}
