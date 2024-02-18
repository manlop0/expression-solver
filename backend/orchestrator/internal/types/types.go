package types

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Expression struct {
	Id          int       `json:"id"`
	Value       string    `json:"value"`
	CreatedDate time.Time `json:"date"`
	Status      int       `json:"status"` //принимает 0,1,2 - в очереди, решается, решен
	Result      string    `json:"result"`
}

type Operations struct {
	Name     string `json:"name"`
	Duration int    `json:"duration"`
}

type Worker struct {
	Id        int    `json:"id"`
	Working   bool   `json:"working"`   //принимает false,true - свободен, решает
	WorkingOn string `json:"workingon"` // выражение над которым работает. Не забудь поменять в database.go || мб не надо
}

type Task struct {
	Exp Expression
	Ops map[string]int
}

type RedisQueue struct {
	Client    *redis.Client
	QueueName string
}

// поместить в очередь
func (rq *RedisQueue) Enqueue(task Task) error {
	jsonData, err := json.Marshal(task)
	if err != nil {
		return err
	}
	ctx := context.Background()
	_, err = rq.Client.LPush(ctx, rq.QueueName, jsonData).Result()
	return err
}
