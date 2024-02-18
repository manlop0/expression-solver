package handlers

import (
	"database/sql"
	"log"
	"orchestrator/internal/types"
	"orchestrator/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AddExpressionHandler(ctx *fiber.Ctx, db *sql.DB, rq types.RedisQueue) error {

	// добавление выражения в бд
	expression := types.Expression{}

	if err := ctx.BodyParser(&expression); err != nil {
		log.Panicf("Error while parsing data: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while parsing data",
		})
	}

	if !utils.ValidateExpression(expression.Value) {
		log.Printf("Invalid expression: %s", expression.Value)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid expression",
		})
	}

	expression.Result = "?"
	expression.Status = 0

	err := db.QueryRow("INSERT INTO expressions (value, date, status, result) VALUES ($1, $2, $3, $4) RETURNING id", expression.Value, expression.CreatedDate, expression.Status, expression.Result).Scan(&expression.Id)
	if err != nil {
		log.Panicf("Error while adding data into db: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while adding data into db",
		})
	}

	operations, err := utils.GetOperationsTime(db)
	if err != nil {
		log.Panicf("Error while getting operations: %s", err)
	}

	task := types.Task{Exp: expression, Ops: operations}

	//отправляем выражение в очередь к агенту
	err = rq.Enqueue(task)
	if err != nil {
		log.Panicf("Error while sending task to agent: %s", err)
	}

	return ctx.JSON(expression)
}

func GetExpressionsHandler(ctx *fiber.Ctx, db *sql.DB) error {
	var exp types.Expression
	var expressions []types.Expression

	rows, err := db.Query("SELECT * FROM expressions")
	if err != nil {
		log.Panicf("Error while getting expressions: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while getting expressions",
		})
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&exp.Id, &exp.Value, &exp.CreatedDate, &exp.Status, &exp.Result); err != nil {
			log.Panicf("Error while scanning rows: %s", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error while scanning rows",
			})
		}
		expressions = append(expressions, exp)
	}

	return ctx.JSON(expressions)
}

func GetOperationsHandler(ctx *fiber.Ctx, db *sql.DB) error {
	var op types.Operations
	var operations []types.Operations

	rows, err := db.Query("SELECT * FROM operations")
	if err != nil {
		log.Panicf("Error while getting operations: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while getting operations",
		})
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&op.Name, &op.Duration); err != nil {
			log.Panicf("Error while scanning rows: %s", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error while scanning rows",
			})
		}
		operations = append(operations, op)
	}

	return ctx.JSON(operations)
}

func ChangeOperationsHandler(ctx *fiber.Ctx, db *sql.DB) error {
	var operations []types.Operations

	if err := ctx.BodyParser(&operations); err != nil {
		log.Panicf("Error while parsing data: %s", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error while parsing data",
		})
	}

	for _, op := range operations {
		_, err := db.Exec("UPDATE operations SET duration=$1 WHERE name=$2", op.Duration, op.Name)
		if err != nil {
			log.Panicf("Error while changing data: %s", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Error while changing data",
			})
		}
	}

	return ctx.JSON(operations)
}

func GetWorkersHandler(ctx *fiber.Ctx, db *sql.DB) error {

	var w types.Worker
	var workers []types.Worker
	rows, err := db.Query("SELECT * FROM workers")
	if err != nil {
		log.Panicf("Error while getting workers: %s", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while getting workers",
		})
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&w.Id, &w.Working, &w.WorkingOn); err != nil {
			log.Panicf("Error while scanning rows: %s", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error while scanning rows",
			})
		}
		workers = append(workers, w)
	}

	return ctx.JSON(workers)
}
