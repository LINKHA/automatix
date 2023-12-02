package servercode

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	SERVER_CODE_KEY_PREFIX   = "server_code:key:"
	SERVER_CODE_VALUE_PREFIX = "server_code:value:"
)

func serverCodeKey(code string) string {
	return fmt.Sprintf("%s%s", SERVER_CODE_KEY_PREFIX, code)
}

func serverCodeValue(userId string, serverId string) string {
	return fmt.Sprintf("%s%s_%s", SERVER_CODE_VALUE_PREFIX, userId, serverId)
}

func GenServerCode(redis *redis.Redis, userId int64, serverId string) string {
	uuidRes, _ := uuid.NewUUID()
	code := uuidRes.String()
	userIdStr := strconv.FormatInt(userId, 10)

	codeKey := serverCodeKey(code)
	redis.Hmset(codeKey, map[string]string{
		"UserId":   userIdStr,
		"ServerId": serverId,
	})

	codeValue := serverCodeValue(userIdStr, serverId)
	redis.Set(codeValue, code)

	redis.Expire(codeKey, 600)
	redis.Expire(codeValue, 600)

	return code
}

func UseServerCode(redis *redis.Redis, code string) (int64, string) {
	userId, serverId := GetServerCode(redis, code)
	redis.Hdel(code)
	return userId, serverId
}

func GetServerCode(redis *redis.Redis, code string) (int64, string) {
	codeKey := serverCodeKey(code)

	serverInfo, err := redis.Hmget(codeKey, "UserId", "ServerId")
	if err != nil {
		return 0, ""
	}
	userId, _ := strconv.ParseInt(serverInfo[0], 10, 64)
	serverId := serverInfo[1]
	return userId, serverId
}

func FindServerCode(redis *redis.Redis, userId int64, serverId string) string {
	userIdStr := strconv.FormatInt(userId, 10)
	codeValue := serverCodeValue(userIdStr, serverId)
	serverCode, err := redis.Get(codeValue)
	if err != nil {
		return ""
	}
	return serverCode
}
