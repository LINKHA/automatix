package svc

import (
	"automatix/common/utils"
	"errors"
)

type StreamManager struct {
	streams utils.ShardLockMaps
}

func NewStreamManager() *StreamManager {
	return &StreamManager{
		streams: utils.NewShardLockMaps(),
	}
}

func (s *StreamManager) Add(key string, stream interface{}) {

	s.streams.Set(key, stream)
}

func (s *StreamManager) Remove(key string) {
	s.streams.Remove(key)
}

func (s *StreamManager) Get(key string) (interface{}, error) {
	if conn, ok := s.streams.Get(key); ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")
}
