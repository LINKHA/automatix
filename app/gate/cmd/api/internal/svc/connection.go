package svc

import (
	"net"
)

type Connection struct {
	conn   net.Conn
	connID string
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
