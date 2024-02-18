package server

import (
	"log"
	"orchestrator/internal/database"
	"orchestrator/internal/orchestrator/handlers"
	"orchestrator/internal/utils"
	"orchestrator/pkg/redispkg"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {
	db := database.ConnectDb()
	defer db.Close()

	rq := redispkg.GetRedisQueue()

	utils.QueueLoosenExpressions(db, rq)

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Response().Header.Set("Access-Control-Allow-Origin", "*")
		c.Response().Header.Set("Access-Control-Allow-Headers", "Content-Type")
		c.Response().Header.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Response().Header.Set("Content-Type", "application/json")

		return c.Next()
	})

	app.Options("/api/addExpression", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/api/addExpression", func(c *fiber.Ctx) error {
		return handlers.AddExpressionHandler(c, db, rq)
	})

	app.Get("/api/getExpressions", func(c *fiber.Ctx) error {
		return handlers.GetExpressionsHandler(c, db)
	})

	app.Get("/api/getOperations", func(c *fiber.Ctx) error {
		return handlers.GetOperationsHandler(c, db)
	})

	app.Options("/api/changeOperations", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	app.Put("/api/changeOperations", func(c *fiber.Ctx) error {
		return handlers.ChangeOperationsHandler(c, db)
	})

	app.Get("/api/getWorkers", func(c *fiber.Ctx) error {
		return handlers.GetWorkersHandler(c, db)
	})

	log.Fatal(app.Listen(":8000"))

}
