package handler

import (
	"automatix/app/gate/cmd/api/internal/logic"
	"automatix/app/gate/cmd/api/internal/svc"
	rolemanagerPb "automatix/app/rolemanager/cmd/rpc/pb"
	roommanagerPb "automatix/app/roommanager/cmd/rpc/pb"
	"context"
	"fmt"
)

func RegisterHandlers(ctx *svc.ServiceContext) {
	//RoleManager rpc
	createRole_Client, createRole_Client_err := ctx.RoleManagerRpc.CreateRoleStream(context.Background())
	if createRole_Client_err == nil {
		createRole_GrpcConn := logic.NewGrpcConnection[
			rolemanagerPb.Rolemanager_CreateRoleStreamClient,
			rolemanagerPb.CreateRoleReq,
			rolemanagerPb.CreateRoleResp,
		](createRole_Client, "")

		logic.NewGrpcConn(&createRole_GrpcConn)

		go createRole_GrpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", createRole_Client_err)
	}

	//RoomManager rpc
	createRoleStream_Client, createRoleStream_Client_err := ctx.RoomManagerRpc.JoinRoomStream(context.Background())
	if createRoleStream_Client_err == nil {
		createRoleStream_GrpcConn := logic.NewGrpcConnection[
			roommanagerPb.Roommanager_JoinRoomStreamClient,
			roommanagerPb.JoinRoomStreamReq,
			roommanagerPb.JoinRoomResp,
		](createRoleStream_Client, "")

		logic.NewGrpcConn(&createRoleStream_GrpcConn)

		go createRoleStream_GrpcConn.Start()
	} else {
		fmt.Println("rpc register err: ", createRoleStream_Client_err)
	}

}
