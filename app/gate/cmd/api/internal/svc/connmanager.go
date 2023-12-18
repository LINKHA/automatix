package svc

import "automatix/common/utils"

type ConnManager struct {
	connections utils.ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: utils.NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(conn Connection) {

	connMgr.connections.Set(conn.GetConnID(), conn)
}

func (connMgr *ConnManager) Remove(conn Connection) {

	connMgr.connections.Remove(conn.GetConnID())
}
