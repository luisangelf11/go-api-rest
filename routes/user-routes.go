package routes

import (
	"api-rest/handlers"

	"github.com/gofiber/fiber/v3"
)

func UserRoutes(router fiber.Router) {
	users := router.Group("/users")

	users.Get("/", handlers.GetUsers)

	users.Get("/:id", handlers.GetUser)

	users.Post("/", handlers.CreateUser)

	users.Delete("/:id", handlers.DeleteUser)
}
