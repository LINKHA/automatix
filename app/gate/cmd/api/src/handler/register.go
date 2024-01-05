package handler

import (
	"context"
	"fmt"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	rolemanagerPb "github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	roommanagerPb "github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
)

func RegisterHandlers(ctx *ServiceContext, s iface.IServer) {

	grpcConnManager, err := s.GetGrpcConnManager()
	if err != nil {
		fmt.Println("server grpc conn manager err: ")
	}

	//RoleManager rpc
	createRole_Client, createRole_Client_err := ctx.RoleManagerRpc.CreateRoleStream(context.Background())
	if createRole_Client_err == nil {

		grpcConn := net.NewGrpcConnection[
			rolemanagerPb.Rolemanager_CreateRoleStreamClient,
			*rolemanagerPb.CreateRoleReq,
			*rolemanagerPb.CreateRoleResp,
		](
			createRole_Client,
			1,
			func() *rolemanagerPb.CreateRoleReq { return &rolemanagerPb.CreateRoleReq{} },
		)

		grpcConnManager.Add(grpcConn)
		go grpcConn.Start()

	} else {
		fmt.Println("rpc register err: ", createRole_Client_err)
	}

	//RoomManager rpc
	createRoleStream_Client, createRoleStream_Client_err := ctx.RoomManagerRpc.JoinRoomStream(context.Background())
	if createRoleStream_Client_err == nil {
		grpcConn := net.NewGrpcConnection[
			roommanagerPb.Roommanager_JoinRoomStreamClient,
			*roommanagerPb.JoinRoomStreamReq,
			*roommanagerPb.JoinRoomResp,
		](
			createRoleStream_Client,
			2,
			func() *roommanagerPb.JoinRoomStreamReq { return &roommanagerPb.JoinRoomStreamReq{} },
		)

		grpcConnManager.Add(grpcConn)

		go grpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", createRoleStream_Client_err)
	}

}
