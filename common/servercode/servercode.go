package servercode

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func GenServerCode(redis *redis.Redis, userId int64, serverId string) string {
	uuidKey, _ := uuid.NewUUID()
	redis.Hmset(uuidKey.String(), map[string]string{
		"UserId":   strconv.FormatInt(userId, 10),
		"ServerId": serverId,
	})
	redis.Expire(uuidKey.String(), 600)
	return uuidKey.String()
}

func UseServerCode(redis *redis.Redis, code string) (int64, string) {
	serverInfo, err := redis.Hmget(code, "UserId", "ServerId")
	if err != nil {
		return 0, ""
	}
	redis.Hdel(code)
	userId, _ := strconv.ParseInt(serverInfo[0], 10, 64)
	serverId := serverInfo[1]

	return userId, serverId
}

func GetServerCode(redis *redis.Redis, code string) (int64, string) {
	serverInfo, err := redis.Hmget(code, "UserId", "ServerId")
	if err != nil {
		return 0, ""
	}
	userId, _ := strconv.ParseInt(serverInfo[0], 10, 64)
	serverId := serverInfo[1]
	return userId, serverId
}
