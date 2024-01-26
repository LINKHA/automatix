package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"google.golang.org/protobuf/proto"

	rolemanagerPb "github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	roommanagerPb "github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
)

func RegisterHandlers(ctx *ServiceContext, s iface.IServer) {
	grpcConnManager, err := s.GetGrpcConnManager()
	if err != nil {
		fmt.Println("server grpc conn manager err: ")
	}
	RegisterRoleManager(grpcConnManager, ctx, s)
	RegisterRoomManager(grpcConnManager, ctx, s)
}

func RegisterRoleManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	registerGrpcConn[
		rolemanagerPb.Rolemanager_CreateRoleStreamClient,
		*rolemanagerPb.CreateRoleReq,
		*rolemanagerPb.CreateRoleResp,
	](
		grpcConnManager,
		ctx.RoleManagerRpc.CreateRoleStream,
		101,
		func() *rolemanagerPb.CreateRoleReq { return &rolemanagerPb.CreateRoleReq{} },
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_RegisterRoleClient,
		*rolemanagerPb.RegisterRoleReq,
		*rolemanagerPb.RegisterRoleResp,
	](
		grpcConnManager,
		ctx.RoleManagerRpc.RegisterRole,
		102,
		func() *rolemanagerPb.RegisterRoleReq { return &rolemanagerPb.RegisterRoleReq{} },
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_SetRoleClient,
		*rolemanagerPb.SetRoleReq,
		*rolemanagerPb.SetRoleResp,
	](
		grpcConnManager,
		ctx.RoleManagerRpc.SetRole,
		103,
		func() *rolemanagerPb.SetRoleReq { return &rolemanagerPb.SetRoleReq{} },
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_GetRoleClient,
		*rolemanagerPb.GetRoleReq,
		*rolemanagerPb.GetRoleResp,
	](
		grpcConnManager,
		ctx.RoleManagerRpc.GetRole,
		104,
		func() *rolemanagerPb.GetRoleReq { return &rolemanagerPb.GetRoleReq{} },
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_DeleteRoleClient,
		*rolemanagerPb.DeleteRoleReq,
		*rolemanagerPb.DeleteRoleResp,
	](
		grpcConnManager,
		ctx.RoleManagerRpc.DeleteRole,
		105,
		func() *rolemanagerPb.DeleteRoleReq { return &rolemanagerPb.DeleteRoleReq{} },
	)
}

func RegisterRoomManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	registerGrpcConn[
		roommanagerPb.Roommanager_CreateRoomClient,
		*roommanagerPb.CreateRoomReq,
		*roommanagerPb.CreateRoomResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.CreateRoom,
		201,
		func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_JoinRoomClient,
		*roommanagerPb.JoinRoomReq,
		*roommanagerPb.JoinRoomResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.JoinRoom,
		202,
		func() *roommanagerPb.JoinRoomReq { return &roommanagerPb.JoinRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_LeaveRoomClient,
		*roommanagerPb.LeaveRoomReq,
		*roommanagerPb.LeaveRoomResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.LeaveRoom,
		203,
		func() *roommanagerPb.LeaveRoomReq { return &roommanagerPb.LeaveRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchRoomClient,
		*roommanagerPb.MatchRoomReq,
		*roommanagerPb.MatchRoomResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.MatchRoom,
		204,
		func() *roommanagerPb.MatchRoomReq { return &roommanagerPb.MatchRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_GetRoomInfoClient,
		*roommanagerPb.GetRoomInfoReq,
		*roommanagerPb.GetRoomInfoResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.GetRoomInfo,
		205,
		func() *roommanagerPb.GetRoomInfoReq { return &roommanagerPb.GetRoomInfoReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchFinishClient,
		*roommanagerPb.MatchFinishReq,
		*roommanagerPb.MatchFinishResp,
	](
		grpcConnManager,
		ctx.RoomManagerRpc.MatchFinish,
		206,
		func() *roommanagerPb.MatchFinishReq { return &roommanagerPb.MatchFinishReq{} },
	)
}

func registerGrpcConn[T_Client net.StreamClientInterface, T_Req proto.Message, T_Resp proto.Message](
	grpcConnManager iface.IGrpcConnManager,
	rpcClient interface{},
	id uint64,
	newReq func() T_Req,
) {
	go func() {
		retryInterval := 3 * time.Second
		for {
			client, err := rpcClient.(func(context.Context) (interface{}, error))(context.Background())
			if err != nil {
				grpcConn := net.NewGrpcConnection[T_Client, T_Req, T_Resp](
					client.(T_Client),
					id,
					newReq,
				)

				grpcConnManager.Add(grpcConn)
				go grpcConn.Start()
				break
			} else {
				fmt.Printf("rpc register err: %v\n", err)
				fmt.Println("Retrying in", retryInterval)
				<-time.After(retryInterval)
			}
		}
	}()
}
