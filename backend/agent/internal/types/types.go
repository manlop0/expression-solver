package types

import (
	"agent/internal/utils/expressionsolver"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type Expression struct {
	Id          int
	Value       string
	CreatedDate time.Time
	Status      int //принимает 0,1,2 - в очереди, решается, решен
	Result      string
}

type Worker struct {
	Id        int
	Working   bool   //принимает false,true - свободен, решает
	WorkingOn string // выражение над которым работает
}

// решить выражение
func (w *Worker) SolveExpression(task Task, agent *Agent) (string, int, error) {
	agent.Mu.Lock()
	exp, ops := task.Exp, task.Ops

	//указываем что этот воркер занят
	_, err := agent.Db.Exec("UPDATE workers SET working=$1, workingon=$2 WHERE id=$3", w.Working, w.WorkingOn, w.Id)
	if err != nil {
		return exp.Value, w.Id, err
	}

	//указываем что выражение решается
	_, err = agent.Db.Exec("UPDATE expressions SET status=$1 WHERE id=$2", 1, task.Exp.Id)
	if err != nil {
		return exp.Value, w.Id, err
	}
	agent.Mu.Unlock()

	//считаем выражение
	res := fmt.Sprint(expressionsolver.EvaluateExpression(exp.Value, ops))

	defer func() {
		//указываем что воркер освободился
		agent.Mu.Lock()
		agent.NumOfWorkers--

		w.Working = false
		w.WorkingOn = ""
		_, err = agent.Db.Exec("UPDATE workers SET working=$1, workingon=$2 WHERE id=$3", w.Working, w.WorkingOn, w.Id)
		if err != nil {
			log.Panicf("Error while setting worker working to false: %s", err)
		}
		agent.Mu.Unlock()
	}()

	//обновляем данные
	agent.Mu.Lock()
	_, err = agent.Db.Exec("UPDATE expressions SET result=$1, status=$2 WHERE id=$3", res, 2, task.Exp.Id)
	if err != nil {
		return exp.Value, w.Id, err
	}
	agent.Mu.Unlock()

	return exp.Value, exp.Id, nil
}

type Task struct {
	Exp Expression
	Ops map[string]int
}

type Agent struct {
	Workers      []*Worker
	MaxWorkers   int
	NumOfWorkers int
	Rq           RedisQueue
	Db           *sql.DB
	Mu           sync.Mutex
}

type RedisQueue struct {
	Client    *redis.Client
	QueueName string
}

// вытащить из очереди
func (rq *RedisQueue) Dequeue() (Task, error) {
	var task Task

	ctx := context.Background()
	jsonData, err := rq.Client.RPop(ctx, rq.QueueName).Bytes()
	if jsonData == nil {
		return task, errors.New("no data in RedisQueue")
	}
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		return task, err
	}

	return task, nil
}
