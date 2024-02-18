package utils

import (
	"agent/internal/types"
	"database/sql"
)

// инициализация воркеров
func InitializeWorkers(db *sql.DB, maxWorkers int) ([]*types.Worker, error) {
	workers := make([]*types.Worker, maxWorkers)

	_, err := db.Exec("DELETE FROM workers")
	if err != nil {
		return workers, err
	}

	for i := 1; i <= maxWorkers; i++ {
		_, err := db.Exec("INSERT INTO workers (id, working, workingon) VALUES ($1, $2, $3)", i, false, "")
		if err != nil {
			return workers, err
		}
		workers[i-1] = &types.Worker{Id: i, Working: false, WorkingOn: ""}
	}

	return workers, nil
}
