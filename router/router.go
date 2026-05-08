package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"tast-list.com/database"
	"tast-list.com/handlers"
	"tast-list.com/repositories"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handlers.Hello)

	//Task
	taskRepo := repositories.NewTaskRepository(database.DB)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	tasks := api.Group("/tasks")
	tasks.Post("/", taskHandler.CreateTask)
	tasks.Get("/", taskHandler.GetAll)
	tasks.Get("/:id", taskHandler.GetById)
	tasks.Put("/:id", taskHandler.UpdateTaskTask)
	tasks.Delete("/:id", taskHandler.DeleteTask)
}
