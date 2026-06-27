package routes

import (
	"api-rest/handlers"

	"github.com/gofiber/fiber/v3"
)

func TaskRoutes(router fiber.Router) {
	tasks := router.Group("/tasks")

	tasks.Get("/", handlers.GetTasks)

	tasks.Get("/:id", handlers.GetTask)

	tasks.Post("/", handlers.CreateTask)

	tasks.Delete("/:id", handlers.DeleteTask)

	tasks.Patch("/:id", handlers.UpdateTask)
}
