package utils

import "strconv"

func JoinCacheKey(appName string, id int64, keyName string) string {
	return appName + "/" + strconv.FormatInt((id), 10) + "/" + keyName
}
