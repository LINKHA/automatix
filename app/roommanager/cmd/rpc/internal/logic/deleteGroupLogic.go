package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(stream pb.Roommanager_DeleteGroupServer) error {
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
func (l *DeleteGroupLogic) handlerFunc(stream pb.Roommanager_DeleteGroupServer, req *pb.DeleteGroupReq) {
	err := l.svcCtx.Redis.Pipelined(
		func(pipe redis.Pipeliner) error {
			groupKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_GROUP, req.GroupId)
			groupStr, err := l.svcCtx.Redis.Get(groupKey)
			if err != nil {
				return err
			}

			m := &svc.Group{}
			json.Unmarshal([]byte(groupStr), m)

			for _, roleId := range m.Roles {
				roleKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_ROLE, roleId)
				pipe.Expire(context.Background(), roleKey, 86400*time.Second)
			}
			pipe.Del(context.Background(), groupKey)
			return nil
		},
	)

	if err != nil {
		stream.Send(&pb.DeleteGroupResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}
	stream.Send(&pb.DeleteGroupResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
	})
}
