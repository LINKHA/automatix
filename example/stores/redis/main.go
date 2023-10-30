package main

import (
	"github.com/linkha/automatix/core/stores/redis"
)

func main() {
	store := redis.New("localhost:6379")
	store.Set("wan", "111")
}
