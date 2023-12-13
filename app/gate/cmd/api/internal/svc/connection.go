package svc

import (
	"context"
	"fmt"
	"net"
)

type Connection struct {
	conn   net.Conn
	connID string

	ctx    context.Context
	cancel context.CancelFunc

	msgBuffChan chan []byte
}

func newServerConn(server *Server, conn net.Conn, connID string) Connection {

	// Initialize Conn properties
	c := Connection{
		conn:   conn,
		connID: connID,
	}

	server.GetConnMgr().Add(c)
	return c
}

func (c *Connection) GetConnID() string {
	return c.connID
}

func (c *Connection) Start() {
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

func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) StartReader() {
	defer c.Stop()

	buffer := make([]byte, 1024)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			n, err := c.conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err)
				return
			}

			fmt.Printf("Received TCP message: %s\n", string(buffer[:n]))

		}
	}
}

func (c *Connection) StartWriter() {
	for {
		select {
		case data, ok := <-c.msgBuffChan:
			if ok {
				if err := c.Send(data); err != nil {
					fmt.Printf("Send Buff Data error:, %s Conn Writer exit\n", err)
					break
				}

			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) Send(data []byte) error {
	_, err := c.conn.Write(data)

	if err != nil {
		fmt.Printf("SendMsg err data = %+v, err = %+v\n", data, err)
		return err
	}

	return nil
}
