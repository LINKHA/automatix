package net

import (
	"errors"
	"strconv"

	"github.com/LINKHA/automatix/app/gate/cmd/api/src/net/iface"
	"github.com/LINKHA/automatix/common/log"
	"github.com/LINKHA/automatix/common/utils"
)

type ConnManager struct {
	connections utils.ShardLockMaps
}

func newConnManager() *ConnManager {
	return &ConnManager{
		connections: utils.NewShardLockMaps(),
	}
}

func (connMgr *ConnManager) Add(conn iface.IConnection) {

	connMgr.connections.Set(conn.GetConnIdStr(), conn) // 将conn连接添加到ConnManager中

	log.Ins().InfoF("connection add to ConnManager successfully: conn num = %d", connMgr.Len())
}

func (connMgr *ConnManager) Remove(conn iface.IConnection) {

	connMgr.connections.Remove(conn.GetConnIdStr()) // 删除连接信息

	log.Ins().InfoF("connection Remove ConnID=%d successfully: conn num = %d", conn.GetConnID(), connMgr.Len())
}

func (connMgr *ConnManager) Get(connID uint64) (iface.IConnection, error) {

	strConnId := strconv.FormatUint(connID, 10)
	if conn, ok := connMgr.connections.Get(strConnId); ok {
		return conn.(iface.IConnection), nil
	}

	return nil, errors.New("connection not found")
}

// Get2 It is recommended to use this method to obtain connection instances
func (connMgr *ConnManager) Get2(strConnId string) (iface.IConnection, error) {

	if conn, ok := connMgr.connections.Get(strConnId); ok {
		return conn.(iface.IConnection), nil
	}

	return nil, errors.New("connection not found")
}

func (connMgr *ConnManager) Len() int {

	length := connMgr.connections.Count()

	return length
}

func (connMgr *ConnManager) ClearConn() {

	// Stop and delete all connection information
	for item := range connMgr.connections.IterBuffered() {
		val := item.Val
		if conn, ok := val.(iface.IConnection); ok {
			// stop will eventually trigger the deletion of the connection,
			// no additional deletion is required
			conn.Stop()
		}
	}

	log.Ins().InfoF("Clear All Connections successfully: conn num = %d", connMgr.Len())
}

func (connMgr *ConnManager) GetAllConnID() []uint64 {

	strConnIdList := connMgr.connections.Keys()
	ids := make([]uint64, 0, len(strConnIdList))

	for _, strId := range strConnIdList {
		connId, err := strconv.ParseUint(strId, 10, 64)
		if err == nil {
			ids = append(ids, connId)
		} else {
			log.Ins().InfoF("GetAllConnID Id: %d, error: %v", connId, err)
		}
	}

	return ids
}

func (connMgr *ConnManager) GetAllConnIdStr() []string {
	return connMgr.connections.Keys()
}

func (connMgr *ConnManager) Range(cb func(uint64, iface.IConnection, interface{}) error, args interface{}) (err error) {

	connMgr.connections.IterCb(func(key string, v interface{}) {
		conn, _ := v.(iface.IConnection)
		connId, _ := strconv.ParseUint(key, 10, 64)
		err = cb(connId, conn, args)
		if err != nil {
			log.Ins().InfoF("Range key: %v, v: %v, error: %v", key, v, err)
		}
	})

	return err
}

// Range2 It is recommended to use this method to 'Range'
func (connMgr *ConnManager) Range2(cb func(string, iface.IConnection, interface{}) error, args interface{}) (err error) {

	connMgr.connections.IterCb(func(key string, v interface{}) {
		conn, _ := v.(iface.IConnection)
		err = cb(conn.GetConnIdStr(), conn, args)
		if err != nil {
			log.Ins().InfoF("Range2 key: %v, v: %v, error: %v", key, v, err)
		}
	})

	return err
}
