package net

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"google.golang.org/grpc"
)

type StreamClientInterface interface {
	grpc.ClientStream
}

type GrpcConnection[T1 StreamClientInterface, T2 any, T3 any] struct {
	conn      T1
	connId    uint64
	connIdStr string

	ctx    context.Context
	cancel context.CancelFunc

	msgReqChan  chan T2
	msgRespChan chan T3
}

func NewGrpcConnection[T1 StreamClientInterface, T2 any, T3 any](conn T1, connId uint64) iface.IGrpcConnection {

	// Initialize Conn properties
	c := &GrpcConnection[T1, T2, T3]{
		conn:        conn,
		connId:      connId,
		connIdStr:   strconv.FormatUint(connId, 10),
		msgRespChan: make(chan T3, 1000),
	}
	return c
}

func (c *GrpcConnection[T1, T2, T3]) GetConnId() uint64 {
	return c.connId
}

func (c *GrpcConnection[T1, T2, T3]) GetConnIdStr() string {
	return c.connIdStr
}

func (c *GrpcConnection[T1, T2, T3]) Start() {
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

func (c *GrpcConnection[T1, T2, T3]) Stop() {
	c.cancel()
}

func (c *GrpcConnection[T1, T2, T3]) Send(_m interface{}) error {
	m := _m.(T2)

	return c.conn.SendMsg(m)
}

func (c *GrpcConnection[T1, T2, T3]) SendToQueue(_m interface{}) error {
	m := _m.(T2)

	if c.msgReqChan == nil {
		c.msgReqChan = make(chan T2, 1000)
		go c.StartWriter()
	}
	c.msgReqChan <- m

	return nil
}

func (c *GrpcConnection[T1, T2, T3]) StartReader() {
	m := new(T3)
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			err := c.conn.RecvMsg(m)

			if err == io.EOF {
				fmt.Printf("Received GRPC EOF\n")
			}

			if err != nil {
				fmt.Printf("Received GRPC err: %s\n", err)
			}

			fmt.Printf("Received GRPC message: %v\n", m)
		}
	}
}

func (c *GrpcConnection[T1, T2, T3]) StartWriter() {
	for {
		select {
		case data, ok := <-c.msgReqChan:
			if ok {
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
