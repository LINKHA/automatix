package server

import (
	"context"
	"fmt"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/handler"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	commonPb "github.com/LINKHA/automatix/common/proto"
	"google.golang.org/protobuf/proto"

	rolemanagerPb "github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	roommanagerPb "github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
)

func RegisterHandlers(ctx *handler.ServiceContext, s iface.IServer) {
	RegisterRoleManager(ctx, s)
	RegisterRoomManager(ctx, s)
}

func RegisterRoleManager(ctx *handler.ServiceContext, s iface.IServer) {
	registerGrpcConn[
		rolemanagerPb.Rolemanager_RegisterRoleClient,
		*rolemanagerPb.RegisterRoleReq,
		*rolemanagerPb.RegisterRoleResp,
	](
		101,
		s,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_RegisterRoleClient, error) {
			return ctx.RoleManagerRpc.RegisterRole(context.Background())
		},
		func() *rolemanagerPb.RegisterRoleReq { return &rolemanagerPb.RegisterRoleReq{} },
		func() *rolemanagerPb.RegisterRoleResp { return &rolemanagerPb.RegisterRoleResp{} },
		func(connId uint64, req *rolemanagerPb.RegisterRoleReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *rolemanagerPb.RegisterRoleResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_SetRoleClient,
		*rolemanagerPb.SetRoleReq,
		*rolemanagerPb.SetRoleResp,
	](
		102,
		s,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_SetRoleClient, error) {
			return ctx.RoleManagerRpc.SetRole(context.Background())
		},
		func() *rolemanagerPb.SetRoleReq { return &rolemanagerPb.SetRoleReq{} },
		func() *rolemanagerPb.SetRoleResp { return &rolemanagerPb.SetRoleResp{} },
		func(connId uint64, req *rolemanagerPb.SetRoleReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *rolemanagerPb.SetRoleResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_GetRoleClient,
		*rolemanagerPb.GetRoleReq,
		*rolemanagerPb.GetRoleResp,
	](
		103,
		s,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_GetRoleClient, error) {
			return ctx.RoleManagerRpc.GetRole(context.Background())
		},
		func() *rolemanagerPb.GetRoleReq { return &rolemanagerPb.GetRoleReq{} },
		func() *rolemanagerPb.GetRoleResp { return &rolemanagerPb.GetRoleResp{} },
		func(connId uint64, req *rolemanagerPb.GetRoleReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *rolemanagerPb.GetRoleResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		rolemanagerPb.Rolemanager_DeleteRoleClient,
		*rolemanagerPb.DeleteRoleReq,
		*rolemanagerPb.DeleteRoleResp,
	](
		104,
		s,
		func(mctx context.Context) (rolemanagerPb.Rolemanager_DeleteRoleClient, error) {
			return ctx.RoleManagerRpc.DeleteRole(context.Background())
		},
		func() *rolemanagerPb.DeleteRoleReq { return &rolemanagerPb.DeleteRoleReq{} },
		func() *rolemanagerPb.DeleteRoleResp { return &rolemanagerPb.DeleteRoleResp{} },
		func(connId uint64, req *rolemanagerPb.DeleteRoleReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *rolemanagerPb.DeleteRoleResp) uint64 {
			return resp.Header.ConnId
		},
	)
}

func RegisterRoomManager(ctx *handler.ServiceContext, s iface.IServer) {
	registerGrpcConn[
		roommanagerPb.Roommanager_CreateGroupClient,
		*roommanagerPb.CreateGroupReq,
		*roommanagerPb.CreateGroupResp,
	](
		201,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_CreateGroupClient, error) {
			return ctx.RoomManagerRpc.CreateGroup(context.Background())
		},
		func() *roommanagerPb.CreateGroupReq { return &roommanagerPb.CreateGroupReq{} },
		func() *roommanagerPb.CreateGroupResp { return &roommanagerPb.CreateGroupResp{} },
		func(connId uint64, req *roommanagerPb.CreateGroupReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.CreateGroupResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_DeleteGroupClient,
		*roommanagerPb.DeleteGroupReq,
		*roommanagerPb.DeleteGroupResp,
	](
		202,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_DeleteGroupClient, error) {
			return ctx.RoomManagerRpc.DeleteGroup(context.Background())
		},
		func() *roommanagerPb.DeleteGroupReq { return &roommanagerPb.DeleteGroupReq{} },
		func() *roommanagerPb.DeleteGroupResp { return &roommanagerPb.DeleteGroupResp{} },
		func(connId uint64, req *roommanagerPb.DeleteGroupReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.DeleteGroupResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_GetGroupClient,
		*roommanagerPb.GetGroupReq,
		*roommanagerPb.GetGroupResp,
	](
		203,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_GetGroupClient, error) {
			return ctx.RoomManagerRpc.GetGroup(context.Background())
		},
		func() *roommanagerPb.GetGroupReq { return &roommanagerPb.GetGroupReq{} },
		func() *roommanagerPb.GetGroupResp { return &roommanagerPb.GetGroupResp{} },
		func(connId uint64, req *roommanagerPb.GetGroupReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.GetGroupResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_JoinGroupClient,
		*roommanagerPb.JoinGroupReq,
		*roommanagerPb.JoinGroupResp,
	](
		204,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_JoinGroupClient, error) {
			return ctx.RoomManagerRpc.JoinGroup(context.Background())
		},
		func() *roommanagerPb.JoinGroupReq { return &roommanagerPb.JoinGroupReq{} },
		func() *roommanagerPb.JoinGroupResp { return &roommanagerPb.JoinGroupResp{} },
		func(connId uint64, req *roommanagerPb.JoinGroupReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.JoinGroupResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_LeaveGroupClient,
		*roommanagerPb.LeaveGroupReq,
		*roommanagerPb.LeaveGroupResp,
	](
		205,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_LeaveGroupClient, error) {
			return ctx.RoomManagerRpc.LeaveGroup(context.Background())
		},
		func() *roommanagerPb.LeaveGroupReq { return &roommanagerPb.LeaveGroupReq{} },
		func() *roommanagerPb.LeaveGroupResp { return &roommanagerPb.LeaveGroupResp{} },
		func(connId uint64, req *roommanagerPb.LeaveGroupReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.LeaveGroupResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_CreateRoomClient,
		*roommanagerPb.CreateRoomReq,
		*roommanagerPb.CreateRoomResp,
	](
		206,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_CreateRoomClient, error) {
			return ctx.RoomManagerRpc.CreateRoom(context.Background())
		},
		func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
		func() *roommanagerPb.CreateRoomResp { return &roommanagerPb.CreateRoomResp{} },
		func(connId uint64, req *roommanagerPb.CreateRoomReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.CreateRoomResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_GetRoomClient,
		*roommanagerPb.GetRoomReq,
		*roommanagerPb.GetRoomResp,
	](
		207,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_GetRoomClient, error) {
			return ctx.RoomManagerRpc.GetRoom(context.Background())
		},
		func() *roommanagerPb.GetRoomReq { return &roommanagerPb.GetRoomReq{} },
		func() *roommanagerPb.GetRoomResp { return &roommanagerPb.GetRoomResp{} },
		func(connId uint64, req *roommanagerPb.GetRoomReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.GetRoomResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchRoomClient,
		*roommanagerPb.MatchRoomReq,
		*roommanagerPb.MatchRoomResp,
	](
		208,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_MatchRoomClient, error) {
			return ctx.RoomManagerRpc.MatchRoom(context.Background())
		},
		func() *roommanagerPb.MatchRoomReq { return &roommanagerPb.MatchRoomReq{} },
		func() *roommanagerPb.MatchRoomResp { return &roommanagerPb.MatchRoomResp{} },
		func(connId uint64, req *roommanagerPb.MatchRoomReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.MatchRoomResp) uint64 {
			return resp.Header.ConnId
		},
	)

	registerGrpcConn[
		roommanagerPb.Roommanager_MatchFinishClient,
		*roommanagerPb.MatchFinishReq,
		*roommanagerPb.MatchFinishResp,
	](
		209,
		s,
		func(mctx context.Context) (roommanagerPb.Roommanager_MatchFinishClient, error) {
			return ctx.RoomManagerRpc.MatchFinish(context.Background())
		},
		func() *roommanagerPb.MatchFinishReq { return &roommanagerPb.MatchFinishReq{} },
		func() *roommanagerPb.MatchFinishResp { return &roommanagerPb.MatchFinishResp{} },
		func(connId uint64, req *roommanagerPb.MatchFinishReq) {
			req.Header = &commonPb.Header{
				ConnId: connId,
			}
		},
		func(resp *roommanagerPb.MatchFinishResp) uint64 {
			return resp.Header.ConnId
		},
	)
}

func registerGrpcConn[T_Client net.StreamClientInterface, T_Req proto.Message,
	T_Resp proto.Message](
	id uint64,
	s iface.IServer,
	rpcClient func(mctx context.Context) (T_Client, error),
	newReq func() T_Req,
	newResp func() T_Resp,
	beforeHook func(uint64, T_Req),
	afterHook func(T_Resp) uint64,
) {
	go func() {
		retryInterval := 3 * time.Second
		for {
			client, err := rpcClient(context.Background())
			if err == nil {
				grpcConn := net.NewGrpcConnection[T_Client, T_Req, T_Resp](
					s,
					client,
					id,
					newReq,
					newResp,
					beforeHook,
					afterHook,
				)
				grpcConnManager, _ := s.GetGrpcConnManager()
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
