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

type GetRoomLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoomLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoomLogic {
	return &GetRoomLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoomLogic) GetRoom(stream pb.Roommanager_GetRoomServer) error {
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

func (l *GetRoomLogic) handlerFunc(stream pb.Roommanager_GetRoomServer, req *pb.GetRoomReq) {
	roomKey := fmt.Sprintf("%s:%s", svc.ROOMMANAGER_ROOM, req.RoomId)
	roomStr, err := l.svcCtx.Redis.Get(roomKey)
	if err != nil {
		stream.Send(&pb.GetRoomResp{
			Header:     req.Header,
			ReturnCode: int64(xerr.SERVER_COMMON_ERROR),
		})
		return
	}

	m := &svc.Room{}
	json.Unmarshal([]byte(roomStr), m)

	stream.Send(&pb.GetRoomResp{
		Header:     req.Header,
		ReturnCode: int64(xerr.OK),
		RoomId:     m.RoomId,
		RoomName:   m.RoomName,
		MaxPlayer:  int64(m.MaxPlayer),
		RoleIds:    m.Roles,
	})
}
