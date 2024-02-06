package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	mNet "github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/app/roommanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/log"
	"google.golang.org/protobuf/proto"
)

type CreateGroupRouter struct {
	mNet.BaseRouter
}

func (this *CreateGroupRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.CreateGroupResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestCreateGroup(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.CreateGroupReq{
			RoleId:    "1",
			GroupName: "1",
			MaxPlayer: 5,
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(201, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(201, &CreateGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 2)
}
