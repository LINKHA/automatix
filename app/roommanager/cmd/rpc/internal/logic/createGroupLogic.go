package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(stream pb.Roommanager_CreateGroupServer) error {
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

func (l *CreateGroupLogic) handlerFunc(stream pb.Roommanager_CreateGroupServer, req *pb.CreateGroupReq) {
	groupId := strconv.FormatInt(int64(l.svcCtx.Snowflake.Generate()), 10)
	groupKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_GROUP, groupId)

	// group redis info
	group := &svc.Group{
		GroupID:   groupId,
		GroupName: req.GroupName,
		MaxPlayer: int32(req.MaxPlayer),
		RoomID:    "",
		Roles:     []string{req.RoleId},
	}

	groupJSON, err := json.Marshal(group)
	if err != nil {
		fmt.Println("Error:", err)
		stream.Send(&pb.CreateGroupResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
		return
	}

	l.svcCtx.Redis.Setex(groupKey, string(groupJSON), 86400)

	// role redis info
	roleKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_ROLE, req.RoleId)
	err2 := l.svcCtx.Redis.Hmset(roleKey, map[string]string{
		"RoleId":  req.RoleId,
		"GroupId": groupId,
	})

	l.svcCtx.Redis.Expire(roleKey, 86400)
	if err2 != nil {
		fmt.Println("Error:", err2)
		stream.Send(&pb.CreateGroupResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
		return
	}

	stream.Send(&pb.CreateGroupResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
		GroupId:    groupId,
	})
}
