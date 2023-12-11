package net

import (
	"errors"
	"strconv"
)

type ConnManager struct {
	connections ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(conn Connection) {

	connMgr.connections.Set(conn.GetConnIdStr(), conn)
}

func (connMgr *ConnManager) Remove(conn Connection) {

	connMgr.connections.Remove(conn.GetConnIdStr())
}

func (connMgr *ConnManager) Get(connID uint64) (Connection, error) {

	strConnId := strconv.FormatUint(connID, 10)
	if conn, ok := connMgr.connections.Get(strConnId); ok {
		return conn.(Connection), nil
	}

	return nil, errors.New("connection not found")
}
