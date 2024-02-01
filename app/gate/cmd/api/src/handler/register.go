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
	// RegisterRoomManager(grpcConnManager, ctx, s)
}

func RegisterRoleManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	registerGrpcConn[
		rolemanagerPb.Rolemanager_CreateRoleStreamClient,
		*rolemanagerPb.CreateRoleReq,
		rolemanagerPb.CreateRoleResp,
	](
		101,
		grpcConnManager,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_CreateRoleStreamClient, error) {
			return ctx.RoleManagerRpc.CreateRoleStream(context.Background())
		},
		func() *rolemanagerPb.CreateRoleReq { return &rolemanagerPb.CreateRoleReq{} },
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_RegisterRoleClient,
		*rolemanagerPb.RegisterRoleReq,
		rolemanagerPb.RegisterRoleResp,
	](
		102,
		grpcConnManager,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_RegisterRoleClient, error) {
			return ctx.RoleManagerRpc.RegisterRole(context.Background())
		},
		func() *rolemanagerPb.RegisterRoleReq { return &rolemanagerPb.RegisterRoleReq{} },
	)

	// registerGrpcConn[
	// 	rolemanagerPb.Rolemanager_SetRoleClient,
	// 	*rolemanagerPb.SetRoleReq,
	// 	rolemanagerPb.SetRoleResp,
	// ](
	// 	103,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (rolemanagerPb.Rolemanager_SetRoleClient, error) {
	// 		return ctx.RoleManagerRpc.SetRole(context.Background())
	// 	},
	// 	func() *rolemanagerPb.SetRoleReq { return &rolemanagerPb.SetRoleReq{} },
	// )

	// registerGrpcConn[
	// 	rolemanagerPb.Rolemanager_GetRoleClient,
	// 	*rolemanagerPb.GetRoleReq,
	// 	rolemanagerPb.GetRoleResp,
	// ](
	// 	104,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (rolemanagerPb.Rolemanager_GetRoleClient, error) {
	// 		return ctx.RoleManagerRpc.GetRole(context.Background())
	// 	},
	// 	func() *rolemanagerPb.GetRoleReq { return &rolemanagerPb.GetRoleReq{} },
	// )

	// registerGrpcConn[
	// 	rolemanagerPb.Rolemanager_DeleteRoleClient,
	// 	*rolemanagerPb.DeleteRoleReq,
	// 	rolemanagerPb.DeleteRoleResp,
	// ](
	// 	105,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (rolemanagerPb.Rolemanager_DeleteRoleClient, error) {
	// 		return ctx.RoleManagerRpc.DeleteRole(context.Background())
	// 	},
	// 	func() *rolemanagerPb.DeleteRoleReq { return &rolemanagerPb.DeleteRoleReq{} },
	// )
}

func RegisterRoomManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	registerGrpcConn[
		roommanagerPb.Roommanager_CreateRoomClient,
		*roommanagerPb.CreateRoomReq,
		roommanagerPb.CreateRoomResp,
	](
		201,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_CreateRoomClient, error) {
			return ctx.RoomManagerRpc.CreateRoom(context.Background())
		},
		func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_JoinRoomClient,
		*roommanagerPb.JoinRoomReq,
		roommanagerPb.JoinRoomResp,
	](
		202,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_JoinRoomClient, error) {
			return ctx.RoomManagerRpc.JoinRoom(context.Background())
		},
		func() *roommanagerPb.JoinRoomReq { return &roommanagerPb.JoinRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_LeaveRoomClient,
		*roommanagerPb.LeaveRoomReq,
		roommanagerPb.LeaveRoomResp,
	](
		203,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_LeaveRoomClient, error) {
			return ctx.RoomManagerRpc.LeaveRoom(context.Background())
		},
		func() *roommanagerPb.LeaveRoomReq { return &roommanagerPb.LeaveRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchRoomClient,
		*roommanagerPb.MatchRoomReq,
		roommanagerPb.MatchRoomResp,
	](
		204,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_MatchRoomClient, error) {
			return ctx.RoomManagerRpc.MatchRoom(context.Background())
		},
		func() *roommanagerPb.MatchRoomReq { return &roommanagerPb.MatchRoomReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_GetRoomInfoClient,
		*roommanagerPb.GetRoomInfoReq,
		roommanagerPb.GetRoomInfoResp,
	](
		205,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_GetRoomInfoClient, error) {
			return ctx.RoomManagerRpc.GetRoomInfo(context.Background())
		},
		func() *roommanagerPb.GetRoomInfoReq { return &roommanagerPb.GetRoomInfoReq{} },
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchFinishClient,
		*roommanagerPb.MatchFinishReq,
		roommanagerPb.MatchFinishResp,
	](
		206,
		grpcConnManager,
		func(mctx context.Context) (roommanagerPb.Roommanager_MatchFinishClient, error) {
			return ctx.RoomManagerRpc.MatchFinish(context.Background())
		},
		func() *roommanagerPb.MatchFinishReq { return &roommanagerPb.MatchFinishReq{} },
	)
}

func registerGrpcConn[T_Client net.StreamClientInterface, T_Req proto.Message, T_Resp any](
	id uint64,
	grpcConnManager iface.IGrpcConnManager,
	rpcClient func(mctx context.Context) (T_Client, error),
	newReq func() T_Req,
) {
	go func() {
		retryInterval := 3 * time.Second
		for {
			client, err := rpcClient(context.Background())
			if err == nil {
				grpcConn := net.NewGrpcConnection[T_Client, T_Req, T_Resp](
					client,
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
