package redispkg

import (
	"orchestrator/internal/types"
	"os"

	"github.com/redis/go-redis/v9"
)

func GetRedisQueue() types.RedisQueue {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	rq := types.RedisQueue{Client: client, QueueName: "tasks"}
	return rq
}
