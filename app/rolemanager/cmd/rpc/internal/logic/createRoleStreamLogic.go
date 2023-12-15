package logic

import (
	"context"
	"fmt"

	"automatix/app/rolemanager/cmd/rpc/internal/svc"
	"automatix/app/rolemanager/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRoleStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleStreamLogic {
	return &CreateRoleStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRoleStreamLogic) CreateRoleStream(stream pb.Rolemanager_CreateRoleStreamServer) error {
	// todo: add your logic here and delete this line
	fmt.Printf("1------------------   :   %d", 111)
	fmt.Println(stream)

	stream.Send(&pb.CreateRoleResp{
		ReturnCode: 2,
	})

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("create role stream\n")
	// 	stream.Send(&pb.CreateRoleResp{
	// 		ReturnCode: 1,
	// 	})

	// 	time.Sleep(time.Duration(1) * time.Second)
	// 	// // 读取服务器的响应
	// 	// buffer := make([]byte, 1024)
	// 	// n, err := conn.Read(buffer)
	// 	// if err != nil {
	// 	// 	fmt.Println("Error reading from server:", err)
	// 	// 	return
	// 	// }

	// 	// receivedData := buffer[:n]
	// 	// fmt.Printf("Received from server: %s\n", receivedData)
	// }
	return nil
}
