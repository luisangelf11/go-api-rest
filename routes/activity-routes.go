package routes

import (
	"api-rest/database"
	"api-rest/models"

	"github.com/gofiber/fiber/v3"
)

func ActivityRoutes(router fiber.Router) {
	activity := router.Group("/activity")

	activity.Get("/", func(c fiber.Ctx) error {
		var activity []models.ActivityLog

		result := database.DB.Find(&activity)

		if result.Error != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"error": "No se han podido obtener los datos de auditoría",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Datos de auditoría",
			"data":    activity,
		})
	})
}
