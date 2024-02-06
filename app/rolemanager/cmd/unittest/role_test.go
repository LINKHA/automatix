package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	mNet "github.com/LINKHA/automatix/app/gate/cmd/api/src/net/net"
	"github.com/LINKHA/automatix/app/rolemanager/cmd/rpc/pb"
	"github.com/LINKHA/automatix/common/log"
	"google.golang.org/protobuf/proto"
)

type RegisterRoleRouter struct {
	mNet.BaseRouter
}

func (this *RegisterRoleRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.RegisterRoleResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestRegisterRole(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.RegisterRoleReq{
			AccountId:     "1",
			ServerId:      "2",
			TemplateValue: "{}",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(101, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(101, &RegisterRoleRouter{})
	client.Start()
	time.Sleep(time.Second * 2)
}

type DeleteRoleRouter struct {
	mNet.BaseRouter
}

func (this *DeleteRoleRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.DeleteRoleResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestDeleteRole(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.DeleteRoleReq{
			RoleId: "1754784763119407104",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(104, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})
	client.AddRouter(104, &DeleteRoleRouter{})
	client.Start()

	time.Sleep(time.Second * 2)
}

type SetRoleRouter struct {
	mNet.BaseRouter
}

func (this *SetRoleRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.SetRoleResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestSetRole(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		tmp := make(map[string]interface{})
		tmp["level"] = 1
		templateValue, _ := json.Marshal(tmp)

		data := &pb.SetRoleReq{
			RoleId:        "1754784763119407104",
			TemplateValue: string(templateValue),
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(102, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})
	client.AddRouter(102, &SetRoleRouter{})
	client.Start()

	time.Sleep(time.Second * 2)
}

type GetRoleRouter struct {
	mNet.BaseRouter
}

func (this *GetRoleRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.GetRoleResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestGetRole(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		go func() {
			data := &pb.GetRoleReq{
				RoleId: "1754784763119407104",
			}
			msg, _ := proto.Marshal(data)

			err := conn.SendMsg(103, msg)
			if err != nil {
				fmt.Println(err)
				log.Error(err)
			}
		}()
	})
	client.AddRouter(103, &GetRoleRouter{})
	client.Start()

	time.Sleep(time.Second * 2)
}
