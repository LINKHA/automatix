package svc

type ConnManager struct {
	connections ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(conn Connection) {

	connMgr.connections.Set(conn.GetConnID(), conn)
}

func (connMgr *ConnManager) Remove(conn Connection) {

	connMgr.connections.Remove(conn.GetConnID())
}
