package utils

import (
	"database/sql"
	"log"
	"orchestrator/internal/types"
)

// функция, проверяющая есть ли в бд выражения, требующие вычисления, и добавляющая их в очередь
func QueueLoosenExpressions(db *sql.DB, rq types.RedisQueue) {

	//узнаем время выполнения операций
	operations, err := GetOperationsTime(db)
	if err != nil {
		log.Printf("Error while getting operations: %s", err)
		return
	}

	var solvingExp []types.Expression
	var queuedExp []types.Expression

	// Выбираем все записи с status=1
	rows, err := db.Query("SELECT id, value, date, status, result FROM expressions WHERE status = 1")
	if err != nil {
		log.Printf("Error while selecting all exp with status 1: %s", err)
		return
	}
	defer rows.Close()

	// Проходим по результатам запроса и добавляем их в список solvingExp
	for rows.Next() {
		var exp types.Expression
		err := rows.Scan(&exp.Id, &exp.Value, &exp.CreatedDate, &exp.Status, &exp.Result)
		if err != nil {
			log.Printf("Error while scanning all exp with status 1: %s", err)
			return
		}
		solvingExp = append(solvingExp, exp)
	}

	//добавляем в очередь
	for _, exp := range solvingExp {
		task := types.Task{Exp: exp, Ops: operations}
		err := rq.Enqueue(task)
		if err != nil {
			log.Printf("Error in EnqueuFunction: %s", err)
		}
	}

	// Выбираем все записи с status=0
	rows, err = db.Query("SELECT id, value, date, status, result FROM expressions WHERE status = 0")
	if err != nil {
		log.Printf("Error while selecting all exp with status 0: %s", err)
		return
	}
	defer rows.Close()

	// Проходим по результатам запроса и добавляем их в список queuedExp
	for rows.Next() {
		var exp types.Expression
		err := rows.Scan(&exp.Id, &exp.Value, &exp.CreatedDate, &exp.Status, &exp.Result)
		if err != nil {
			log.Printf("Error while scanning all exp with status 0: %s", err)
			return
		}
		queuedExp = append(queuedExp, exp)
	}

	//добавляем в очередь
	for _, exp := range queuedExp {
		task := types.Task{Exp: exp, Ops: operations}
		err := rq.Enqueue(task)
		if err != nil {
			log.Printf("Error in EnqueuFunction: %s", err)
		}
	}

}

// узнать время выполнения операций
func GetOperationsTime(db *sql.DB) (map[string]int, error) {

	var op types.Operations
	operations := make(map[string]int)

	rows, err := db.Query("SELECT * FROM operations")
	if err != nil {
		return operations, err
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&op.Name, &op.Duration); err != nil {
			return operations, err
		}
		operations[op.Name] = op.Duration
	}
	return operations, nil
}

// Валидация выражения
func ValidateExpression(expression string) bool {

	if expression == "" {
		return false
	}

	// Проверка на наличие некорректных символов
	for _, char := range expression {
	
		if char != '+' && char != '-' && char != '*' && char != '/' && char != '(' && char != ')' && (char < '0' || char > '9') {
			return false
		}
	}

	// Проверка на корректный порядок символов
	for i := 0; i < len(expression)-1; i++ {
		if (expression[i] == '+' || expression[i] == '-' || expression[i] == '*' || expression[i] == '/') && (expression[i+1] == '+' || expression[i+1] == '-' || expression[i+1] == '*' || expression[i+1] == '/') {
			return false
		}
	}

	// Проверка на наличие хотя бы одного числа
	hasNumber := false
	for _, char := range expression {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}

	return hasNumber
}
