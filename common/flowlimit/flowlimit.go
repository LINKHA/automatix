package flowlimit

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SlidingWindowConf struct {
	Rate       int
	WindowSize int
}

const (
	FLOWLIMIT_MIN_TIME = 0
	FLOWLIMIT_MAX_TIME = 9999999999999
)

/*
rate: The number of requests allowed in the window
windowSize: In milliseconds
*/
func SlidingWindow(redis *redis.Redis, serverId string, rate int64, windowSize int64) bool {
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)
	windowStart := currentTime - int64(windowSize)

	slidingKey := fmt.Sprintf("sliding_window:%s", serverId)
	uuidKey, _ := uuid.NewUUID()

	_, err := redis.Zadd(slidingKey, currentTime, uuidKey.String())
	redis.Zremrangebyscore(slidingKey, FLOWLIMIT_MIN_TIME, windowStart)
	totalCount, err := redis.Zcount(slidingKey, windowStart, FLOWLIMIT_MAX_TIME)

	if err != nil {
		panic(err)
	}

	if rate < int64(totalCount) {
		return false
	} else {
		return true
	}
}

func SlidingWindowCount(redis *redis.Redis, serverId string) int {
	slidingKey := fmt.Sprintf("sliding_window:%s", serverId)
	count, _ := redis.Zcount(slidingKey, FLOWLIMIT_MIN_TIME, FLOWLIMIT_MAX_TIME)
	return count
}
