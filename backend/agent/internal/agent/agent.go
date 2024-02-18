package agentService

import (
	"agent/internal/database"
	"agent/internal/pkg/redispkg"
	"agent/internal/types"
	"agent/internal/utils"

	"log"
	"os"
	"strconv"
)

func StartAgent() {
	db := database.ConnectDb()
	defer db.Close()

	rq := redispkg.GetRedisQueue()

	maxWorkers, err := strconv.Atoi(os.Getenv("NUM_OF_WORKERS"))
	if err != nil {
		log.Fatalf("Error while converting num_of_workers: %s", err)
	}

	workers, err := utils.InitializeWorkers(db, maxWorkers)
	if err != nil {
		log.Fatalf("Error in initializeWorkers function: %s", err)
	}

	//создание и запуск агента
	agent := &types.Agent{Workers: workers, NumOfWorkers: 0, MaxWorkers: maxWorkers, Rq: rq, Db: db}

	for {
		//проверка на свободных воркеров
		if agent.MaxWorkers-agent.NumOfWorkers > 0 {
			//получение таски
			task, err := agent.Rq.Dequeue()
			if err != nil && err.Error() != "no data in RedisQueue" {
				log.Panicf("Error in Dequeu function: %s", err)
			} else if task.Exp.Value != "" {
				for _, w := range agent.Workers {
					if !w.Working {
						w.Working = true
						agent.NumOfWorkers++
						w.WorkingOn = task.Exp.Value
						go func() {
							expression, id, err := w.SolveExpression(task, agent)
							if err != nil {
								log.Panicf("Problems with solving expression '%s' in worker '%d': %s", expression, id, err)
							}
						}()
						break
					}
				}
			}
		}
	}
}
