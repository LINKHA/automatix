package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	commonPb "github.com/LINKHA/automatix/common/proto"
	"google.golang.org/protobuf/proto"

	rolemanagerPb "github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
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
		rolemanagerPb.Rolemanager_RegisterRoleClient,
		*rolemanagerPb.RegisterRoleReq,
		rolemanagerPb.RegisterRoleResp,
	](
		101,
		grpcConnManager,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_RegisterRoleClient, error) {
			return ctx.RoleManagerRpc.RegisterRole(context.Background())
		},
		func() *rolemanagerPb.RegisterRoleReq { return &rolemanagerPb.RegisterRoleReq{} },
		func(connId uint64, req *rolemanagerPb.RegisterRoleReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
	)

	// registerGrpcConn[
	// 	rolemanagerPb.Rolemanager_SetRoleClient,
	// 	*rolemanagerPb.SetRoleReq,
	// 	rolemanagerPb.SetRoleResp,
	// ](
	// 	102,
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
	// 	103,
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
	// 	104,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (rolemanagerPb.Rolemanager_DeleteRoleClient, error) {
	// 		return ctx.RoleManagerRpc.DeleteRole(context.Background())
	// 	},
	// 	func() *rolemanagerPb.DeleteRoleReq { return &rolemanagerPb.DeleteRoleReq{} },
	// )
}

func RegisterRoomManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_CreateGroupClient,
	// 	*roommanagerPb.CreateGroupReq,
	// 	roommanagerPb.CreateGroupResp,
	// ](
	// 	201,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_CreateGroupClient, error) {
	// 		return ctx.RoomManagerRpc.CreateGroup(context.Background())
	// 	},
	// 	func() *roommanagerPb.CreateGroupReq { return &roommanagerPb.CreateGroupReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_DeleteGroupClient,
	// 	*roommanagerPb.DeleteGroupReq,
	// 	roommanagerPb.DeleteGroupResp,
	// ](
	// 	202,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_DeleteGroupClient, error) {
	// 		return ctx.RoomManagerRpc.DeleteGroup(context.Background())
	// 	},
	// 	func() *roommanagerPb.DeleteGroupReq { return &roommanagerPb.DeleteGroupReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_GetGroupClient,
	// 	*roommanagerPb.GetGroupReq,
	// 	roommanagerPb.GetGroupResp,
	// ](
	// 	203,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_GetGroupClient, error) {
	// 		return ctx.RoomManagerRpc.GetGroup(context.Background())
	// 	},
	// 	func() *roommanagerPb.GetGroupReq { return &roommanagerPb.GetGroupReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_JoinGroupClient,
	// 	*roommanagerPb.JoinGroupReq,
	// 	roommanagerPb.JoinGroupResp,
	// ](
	// 	204,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_JoinGroupClient, error) {
	// 		return ctx.RoomManagerRpc.JoinGroup(context.Background())
	// 	},
	// 	func() *roommanagerPb.JoinGroupReq { return &roommanagerPb.JoinGroupReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_LeaveGroupClient,
	// 	*roommanagerPb.LeaveGroupReq,
	// 	roommanagerPb.LeaveGroupResp,
	// ](
	// 	205,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_LeaveGroupClient, error) {
	// 		return ctx.RoomManagerRpc.LeaveGroup(context.Background())
	// 	},
	// 	func() *roommanagerPb.LeaveGroupReq { return &roommanagerPb.LeaveGroupReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_CreateRoomClient,
	// 	*roommanagerPb.CreateRoomReq,
	// 	roommanagerPb.CreateRoomResp,
	// ](
	// 	206,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_CreateRoomClient, error) {
	// 		return ctx.RoomManagerRpc.CreateRoom(context.Background())
	// 	},
	// 	func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_CreateRoomClient,
	// 	*roommanagerPb.CreateRoomReq,
	// 	roommanagerPb.CreateRoomResp,
	// ](
	// 	206,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_CreateRoomClient, error) {
	// 		return ctx.RoomManagerRpc.CreateRoom(context.Background())
	// 	},
	// 	func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_MatchRoomClient,
	// 	*roommanagerPb.MatchRoomReq,
	// 	roommanagerPb.MatchRoomResp,
	// ](
	// 	207,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_MatchRoomClient, error) {
	// 		return ctx.RoomManagerRpc.MatchRoom(context.Background())
	// 	},
	// 	func() *roommanagerPb.MatchRoomReq { return &roommanagerPb.MatchRoomReq{} },
	// )

	// registerGrpcConn[
	// 	roommanagerPb.Roommanager_MatchFinishClient,
	// 	*roommanagerPb.MatchFinishReq,
	// 	roommanagerPb.MatchFinishResp,
	// ](
	// 	208,
	// 	grpcConnManager,
	// 	func(mctx context.Context) (roommanagerPb.Roommanager_MatchFinishClient, error) {
	// 		return ctx.RoomManagerRpc.MatchFinish(context.Background())
	// 	},
	// 	func() *roommanagerPb.MatchFinishReq { return &roommanagerPb.MatchFinishReq{} },
	// )
}

func registerGrpcConn[T_Client net.StreamClientInterface, T_Req proto.Message, T_Resp any](
	id uint64,
	grpcConnManager iface.IGrpcConnManager,
	rpcClient func(mctx context.Context) (T_Client, error),
	newReq func() T_Req,
	beginHook func(uint64, T_Req),
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
					beginHook,
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
