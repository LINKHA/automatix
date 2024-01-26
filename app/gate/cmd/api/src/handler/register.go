package handler

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	rolemanagerPb "github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	roommanagerPb "github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
)

func RegisterRoleManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	//RoleManager rpc
	CreateRole_Client, CreateRole_Client_err := ctx.RoleManagerRpc.CreateRoleStream(context.Background())
	if CreateRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_CreateRoleStreamClient,
			*rolemanagerPb.CreateRoleReq,
			*rolemanagerPb.CreateRoleResp,
		](
			CreateRole_Client,
			101,
			func() *rolemanagerPb.CreateRoleReq { return &rolemanagerPb.CreateRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", CreateRole_Client_err)
	}

	RegisterRole_Client, RegisterRole_Client_err := ctx.RoleManagerRpc.RegisterRole(context.Background())
	if RegisterRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_RegisterRoleClient,
			*rolemanagerPb.RegisterRoleReq,
			*rolemanagerPb.RegisterRoleResp,
		](
			RegisterRole_Client,
			102,
			func() *rolemanagerPb.RegisterRoleReq { return &rolemanagerPb.RegisterRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", RegisterRole_Client_err)
	}

	SetRole_Client, SetRole_Client_err := ctx.RoleManagerRpc.SetRole(context.Background())
	if SetRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_SetRoleClient,
			*rolemanagerPb.SetRoleReq,
			*rolemanagerPb.SetRoleResp,
		](
			SetRole_Client,
			103,
			func() *rolemanagerPb.SetRoleReq { return &rolemanagerPb.SetRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", SetRole_Client_err)
	}

	GetRole_Client, GetRole_Client_err := ctx.RoleManagerRpc.GetRole(context.Background())
	if GetRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_GetRoleClient,
			*rolemanagerPb.GetRoleReq,
			*rolemanagerPb.GetRoleResp,
		](
			GetRole_Client,
			104,
			func() *rolemanagerPb.GetRoleReq { return &rolemanagerPb.GetRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", GetRole_Client_err)
	}

	DeleteRole_Client, DeleteRole_Client_err := ctx.RoleManagerRpc.DeleteRole(context.Background())
	if DeleteRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_DeleteRoleClient,
			*rolemanagerPb.DeleteRoleReq,
			*rolemanagerPb.DeleteRoleResp,
		](
			DeleteRole_Client,
			105,
			func() *rolemanagerPb.DeleteRoleReq { return &rolemanagerPb.DeleteRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", DeleteRole_Client_err)
	}
}

func RegisterRoomManager(grpcConnManager iface.IGrpcConnManager, ctx *ServiceContext, s iface.IServer) {
	//RoomManager rpc
	CreateRoom_Client, CreateRoom_Client_err := ctx.RoomManagerRpc.CreateRoom(context.Background())
	if CreateRoom_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_CreateRoomClient,
			*roommanagerPb.CreateRoomReq,
			*roommanagerPb.CreateRoomResp,
		](
			CreateRoom_Client,
			201,
			func() *roommanagerPb.CreateRoomReq { return &roommanagerPb.CreateRoomReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", CreateRoom_Client_err)
	}

	JoinRoom_Client, JoinRoom_Client_err := ctx.RoomManagerRpc.JoinRoom(context.Background())
	if JoinRoom_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_JoinRoomClient,
			*roommanagerPb.JoinRoomReq,
			*roommanagerPb.JoinRoomResp,
		](
			JoinRoom_Client,
			202,
			func() *roommanagerPb.JoinRoomReq { return &roommanagerPb.JoinRoomReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", JoinRoom_Client_err)
	}

	LeaveRoom_Client, LeaveRoom_Client_err := ctx.RoomManagerRpc.LeaveRoom(context.Background())
	if LeaveRoom_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_LeaveRoomClient,
			*roommanagerPb.LeaveRoomReq,
			*roommanagerPb.LeaveRoomResp,
		](
			LeaveRoom_Client,
			203,
			func() *roommanagerPb.LeaveRoomReq { return &roommanagerPb.LeaveRoomReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", LeaveRoom_Client_err)
	}

	MatchRoom_Client, MatchRoom_Client_err := ctx.RoomManagerRpc.MatchRoom(context.Background())
	if MatchRoom_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_MatchRoomClient,
			*roommanagerPb.MatchRoomReq,
			*roommanagerPb.MatchRoomResp,
		](
			MatchRoom_Client,
			204,
			func() *roommanagerPb.MatchRoomReq { return &roommanagerPb.MatchRoomReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", MatchRoom_Client_err)
	}

	GetRoomInfo_Client, GetRoomInfo_Client_err := ctx.RoomManagerRpc.GetRoomInfo(context.Background())
	if GetRoomInfo_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_GetRoomInfoClient,
			*roommanagerPb.GetRoomInfoReq,
			*roommanagerPb.GetRoomInfoResp,
		](
			GetRoomInfo_Client,
			205,
			func() *roommanagerPb.GetRoomInfoReq { return &roommanagerPb.GetRoomInfoReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", GetRoomInfo_Client_err)
	}

	MatchFinish_Client, MatchFinish_Client_err := ctx.RoomManagerRpc.MatchFinish(context.Background())
	if MatchFinish_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_MatchFinishClient,
			*roommanagerPb.MatchFinishReq,
			*roommanagerPb.MatchFinishResp,
		](
			MatchFinish_Client,
			206,
			func() *roommanagerPb.MatchFinishReq { return &roommanagerPb.MatchFinishReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", MatchFinish_Client_err)
	}
}

func RegisterHandlers(ctx *ServiceContext, s iface.IServer) {

	grpcConnManager, err := s.GetGrpcConnManager()
	if err != nil {
		fmt.Println("server grpc conn manager err: ")
	}
	RegisterRoleManager(grpcConnManager, ctx, s)
	RegisterRoomManager(grpcConnManager, ctx, s)
}
