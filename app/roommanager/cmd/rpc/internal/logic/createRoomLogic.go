package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

func (l *CreateRoomLogic) CreateRoom(stream pb.Roommanager_CreateRoomServer) error {
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

func (l *CreateRoomLogic) handlerFunc(stream pb.Roommanager_CreateRoomServer, req *pb.CreateRoomReq) {

	err := l.svcCtx.Redis.Pipelined(
		func(pipe redis.Pipeliner) error {
			roomId := strconv.FormatInt(int64(l.svcCtx.Snowflake.Generate()), 10)
			roomKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_ROOM, roomId)
			// group redis info
			room := &svc.Room{
				RoomId:    roomId,
				RoomName:  req.RoomName,
				MaxPlayer: int32(req.MaxPlayer),
				Roles:     []string{},
				Groups:    []string{},
			}
			roomJSON, err := json.Marshal(room)
			if err != nil {
				return err
			}
			pipe.SetEX(context.Background(), roomKey, string(roomJSON), 86400*time.Second)
			return nil
		},
	)

	if err != nil {
		stream.Send(&pb.CreateRoomResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
	}
	stream.Send(&pb.CreateRoomResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
	})
}
