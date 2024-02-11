package server

import (
	"errors"
	"fmt"

	"github.com/LINKHA/automatix/common/utils"
)

type RoleManager struct {
	roles   utils.ShardLockMaps
	Id2Role utils.ShardLockMaps
}

func NewRoleManager() *RoleManager {
	return &RoleManager{
		roles:   utils.NewShardLockMaps(),
		Id2Role: utils.NewShardLockMaps(),
	}
}

func (r *RoleManager) Add(roldId string, connId uint32) {
	r.roles.Set(roldId, connId)
	r.Id2Role.Set(fmt.Sprint(connId), roldId)
}

func (r *RoleManager) Remove(roldId string) {
	connId, err := r.Get(roldId)
	if err != nil {
		return
	}

	r.roles.Remove(roldId)
	r.Id2Role.Remove(fmt.Sprint(connId))
}

func (r *RoleManager) RemoveByConnId(connId uint32) {
	roldId, err := r.GetByConnId(connId)
	if err != nil {
		return
	}

	r.roles.Remove(roldId)
	r.Id2Role.Remove(fmt.Sprint(connId))
}

func (r *RoleManager) Get(roldId string) (uint32, error) {

	if connId, ok := r.roles.Get(roldId); ok {
		return connId.(uint32), nil
	}

	return 0, errors.New("rold id not find")
}

func (r *RoleManager) GetByConnId(connId uint32) (string, error) {

	if roldId, ok := r.Id2Role.Get(fmt.Sprint(connId)); ok {
		return roldId.(string), nil
	}

	return "", errors.New("conn id not find")
}

func (r *RoleManager) Len() int {

	length := r.roles.Count()

	return length
}
