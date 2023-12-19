package svc

import "automatix/common/utils"

type StreamManager struct {
	streams utils.ShardLockMaps
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: utils.NewShardLockMaps(),
	}
}

func (s *StreamManager) Add() {
	s.streams.Set("", "")
}

func (s *StreamManager) Remove() {
	s.streams.Remove("")
}
