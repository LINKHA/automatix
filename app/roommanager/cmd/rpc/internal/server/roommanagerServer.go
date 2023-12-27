// Code generated by goctl. DO NOT EDIT.
// Source: roommanager.proto

package server

import (
	"context"

	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/logic"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/internal/svc"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
)

type RoommanagerServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedRoommanagerServer
}

func NewRoommanagerServer(svcCtx *svc.ServiceContext) *RoommanagerServer {
	return &RoommanagerServer{
		svcCtx: svcCtx,
	}
}

func (s *RoommanagerServer) CreateRoom(ctx context.Context, in *pb.CreateRoomReq) (*pb.CreateRoomResp, error) {
	l := logic.NewCreateRoomLogic(ctx, s.svcCtx)
	return l.CreateRoom(in)
}

func (s *RoommanagerServer) JoinRoom(ctx context.Context, in *pb.JoinRoomReq) (*pb.JoinRoomResp, error) {
	l := logic.NewJoinRoomLogic(ctx, s.svcCtx)
	return l.JoinRoom(in)
}

func (s *RoommanagerServer) JoinRoomStream(stream pb.Roommanager_JoinRoomStreamServer) error {
	l := logic.NewJoinRoomStreamLogic(stream.Context(), s.svcCtx)
	return l.JoinRoomStream(stream)
}
