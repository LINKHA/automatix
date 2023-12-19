package svc

import (
	"automatix/common/utils"
	"errors"
)

type ConnManager struct {
	connections utils.ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: utils.NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(conn IConnection) {

	connMgr.connections.Set(conn.GetConnID(), conn)
}

func (connMgr *ConnManager) Remove(conn IConnection) {

	connMgr.connections.Remove(conn.GetConnID())
}

func (connMgr *ConnManager) Get(strConnId string) (IConnection, error) {

	if conn, ok := connMgr.connections.Get(strConnId); ok {
		return conn.(IConnection), nil
	}
	return nil, errors.New("connection not found")
}
