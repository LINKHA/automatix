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
			RoleId:    "1755109290944761856",
			GroupName: "group_1",
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
	time.Sleep(time.Second * 1)
}

type DeleteGroupRouter struct {
	mNet.BaseRouter
}

func (this *DeleteGroupRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.DeleteGroupResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestDeleteGroup(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.DeleteGroupReq{
			GroupId: "1755120020779700224",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(202, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(202, &DeleteGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}

type GetGroupRouter struct {
	mNet.BaseRouter
}

func (this *GetGroupRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.GetGroupResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestGetGroup(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.GetGroupReq{
			GroupId: "1755116474122375168",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(203, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(203, &GetGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}

type JoinGroupRouter struct {
	mNet.BaseRouter
}

func (this *JoinGroupRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.JoinGroupResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestJoinGroup(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.JoinGroupReq{
			GroupId: "1756283090377904128",
			RoleId:  "1755116474122375169",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(204, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(204, &JoinGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}

type LeaveGroupRouter struct {
	mNet.BaseRouter
}

func (this *LeaveGroupRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.LeaveGroupResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestLeaveGroup(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.LeaveGroupReq{
			GroupId: "1756283090377904128",
			RoleId:  "1755109290944761856",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(205, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(205, &JoinGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}

type CreateRoomRouter struct {
	mNet.BaseRouter
}

func (this *CreateRoomRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.CreateRoomResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestCreateRoom(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.CreateRoomReq{
			RoomName:  "roomName1",
			MaxPlayer: 10,
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(206, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(206, &JoinGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}

type GetRoomRouter struct {
	mNet.BaseRouter
}

func (this *GetRoomRouter) Handle(request iface.IRequest) {
	pbMsg := &pb.GetRoomResp{}
	proto.Unmarshal(request.GetData(), pbMsg)
	fmt.Println("recv from server : msgId=", request.GetMsgID(), ", data=", pbMsg)
}

func TestGetRoom(t *testing.T) {
	client := mNet.NewClient("127.0.0.1", 8999)
	client.SetOnConnStart(func(conn iface.IConnection) {
		data := &pb.GetRoomReq{
			RoomId: "1756338735512817664",
		}
		msg, _ := proto.Marshal(data)

		err := conn.SendMsg(207, msg)
		if err != nil {
			fmt.Println(err)
			log.Error(err)
		}
	})

	client.AddRouter(207, &JoinGroupRouter{})
	client.Start()
	time.Sleep(time.Second * 1)
}
