package routes

import "github.com/gofiber/fiber/v3"

func IndexRoute(router fiber.Router) {
	router.Get("/", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello world",
		})
	})
}
