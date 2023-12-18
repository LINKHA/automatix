package svc

import "automatix/common/utils"

type StreamManager struct {
	streams utils.ShardLockMaps
}

func newStreamManager() *StreamManager {
	return &StreamManager{
		streams: utils.NewShardLockMaps(),
	}
}

func (s *StreamManager) Add() {

}
