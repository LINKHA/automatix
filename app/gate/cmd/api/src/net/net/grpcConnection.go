package net

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type StreamClientInterface interface {
	grpc.ClientStream
}

type ReqChanInfo struct {
	connId uint64
	msg    []byte
}

type GrpcConnection[T_Client StreamClientInterface, T_Req proto.Message, T_Resp proto.Message] struct {
	conn      T_Client
	connId    uint64
	connIdStr string

	server iface.IServer

	ctx    context.Context
	cancel context.CancelFunc

	msgReqChan  chan iface.IRequest
	msgRespChan chan T_Resp

	newT_ReqFunc  func() T_Req
	newT_RespFunc func() T_Resp

	beforeHook func(uint64, T_Req)
	afterHook  func(T_Resp) uint64
}

func NewGrpcConnection[T_Client StreamClientInterface,
	T_Req proto.Message,
	T_Resp proto.Message](server iface.IServer, conn T_Client, connId uint64, newT_ReqFunc func() T_Req, newT_RespFunc func() T_Resp, beforeHook func(uint64, T_Req), afterHook func(T_Resp) uint64) iface.IGrpcConnection {

	// Initialize Conn properties
	c := &GrpcConnection[T_Client, T_Req, T_Resp]{
		conn:          conn,
		connId:        connId,
		connIdStr:     strconv.FormatUint(connId, 10),
		server:        server,
		msgRespChan:   make(chan T_Resp, 1000),
		newT_ReqFunc:  newT_ReqFunc,
		newT_RespFunc: newT_RespFunc,
		beforeHook:    beforeHook,
		afterHook:     afterHook,
	}
	return c
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) GetConnId() uint64 {
	return c.connId
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) GetConnIdStr() string {
	return c.connIdStr
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) Start() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Connection Start() error: %v\n", err)
		}
	}()
	c.ctx, c.cancel = context.WithCancel(context.Background())

	go c.StartReader()

	select {
	case <-c.ctx.Done():
		return
	}
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) Stop() {
	c.cancel()
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) Send(_m interface{}) error {
	m := _m.(T_Req)

	return c.conn.SendMsg(m)
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) SendToReqQueue(request iface.IRequest) error {
	if c.msgReqChan == nil {
		c.msgReqChan = make(chan iface.IRequest, 1000)
		go c.StartWriter()
	}
	c.msgReqChan <- request

	return nil
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) StartReader() {
	m := c.newT_RespFunc()
	go c.HandleResp()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			err := c.conn.RecvMsg(m)
			c.msgRespChan <- m

			if err == io.EOF {
				fmt.Printf("Received GRPC EOF\n")
				return
			}
			fmt.Printf("Received GRPC message: %v\n", m)
			if err != nil {
				fmt.Printf("Received GRPC err: %s\n", err)
				return
			}
		}
	}
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) StartWriter() {
	for {
		select {
		case request, ok := <-c.msgReqChan:
			if ok {
				pbMsg := c.newT_ReqFunc()
				proto.Unmarshal(request.GetData(), pbMsg)
				c.beforeHook(request.GetConnection().GetConnID(), pbMsg)

				if err := c.conn.SendMsg(pbMsg); err != nil {
					fmt.Printf("Send Buff Data error:, %s Conn Writer exit\n", err)
					break
				}

			} else {
				fmt.Println("msgReqChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) HandleResp() {
	for {
		select {
		case resp, ok := <-c.msgRespChan:
			if ok {
				connManager, _ := c.server.GetConnManager()
				connId := c.afterHook(resp)
				conn, err := connManager.Get(connId)
				if err == nil {
					m, _ := proto.Marshal(resp)
					conn.SendBuffMsg(uint32(c.connId), m)
				}
			} else {
				fmt.Println("msgRespChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}
