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

type GrpcConnection[T_Client StreamClientInterface, T_Req proto.Message, T_Resp any] struct {
	conn      T_Client
	connId    uint64
	connIdStr string

	ctx    context.Context
	cancel context.CancelFunc

	msgReqChan  chan T_Req
	msgRespChan chan T_Resp

	newT_ReqFunc func() T_Req

	beginHook func(uint64, T_Req)
}

func NewGrpcConnection[T_Client StreamClientInterface,
	T_Req proto.Message,
	T_Resp any](conn T_Client, connId uint64, newT_ReqFunc func() T_Req, beginHook func(uint64, T_Req)) iface.IGrpcConnection {

	// Initialize Conn properties
	c := &GrpcConnection[T_Client, T_Req, T_Resp]{
		conn:         conn,
		connId:       connId,
		connIdStr:    strconv.FormatUint(connId, 10),
		msgRespChan:  make(chan T_Resp, 1000),
		newT_ReqFunc: newT_ReqFunc,
		beginHook:    beginHook,
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

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) SendToReqQueue(connId uint64, m []byte) error {
	pbMsg := c.newT_ReqFunc()
	proto.Unmarshal(m, pbMsg)
	if c.msgReqChan == nil {
		c.msgReqChan = make(chan T_Req, 1000)
		go c.StartWriter()
	}
	c.msgReqChan <- pbMsg

	return nil
}

func (c *GrpcConnection[T_Client, T_Req, T_Resp]) StartReader() {
	m := new(T_Resp)
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			err := c.conn.RecvMsg(m)
			c.msgRespChan <- *m

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
		case data, ok := <-c.msgReqChan:
			if ok {
				c.beginHook(100, data)
				if err := c.conn.SendMsg(data); err != nil {
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
