package distributedlock

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	DISTRIBUTED_LOCK_KEY = "distributedlock"
)

func AcquireLock(redisCli *redis.Redis, lockKey string, identifier string, expiration time.Duration) (bool, error) {
	err := redisCli.Pipelined(
		func(pipe redis.Pipeliner) error {
			distributedlockKey := fmt.Sprintf("%s:%s", DISTRIBUTED_LOCK_KEY, lockKey)

			pipe.SetNX(context.Background(), distributedlockKey, identifier, expiration)
			return nil
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

func ReleaseLock(redisCli *redis.Redis, lockKey string, identifier string) (bool, error) {
	distributedlockKey := fmt.Sprintf("%s:%s", DISTRIBUTED_LOCK_KEY, lockKey)
	// GET command: Get the current value of the lock key
	currentIdentifier, err := redisCli.Get(distributedlockKey)
	if err != nil {
		return false, err
	}

	// Check if the current identifier matches the provided identifier
	if currentIdentifier == identifier {
		// DEL command: Delete the lock key
		_, err := redisCli.Del(distributedlockKey)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}
