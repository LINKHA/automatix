package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"

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
	groupKey := fmt.Sprintf("roommanager:group:%s", groupId)
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
		return
	}

	l.svcCtx.Redis.Set(groupKey, string(groupJSON))
	fmt.Print("0------------------   :   ", 1)
	// ff, _ := l.svcCtx.Redis.Get(groupKey)
	// fmt.Print("1------------------   :   ", ff)
}
