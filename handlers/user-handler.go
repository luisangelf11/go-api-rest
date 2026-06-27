package handlers

import (
	"api-rest/database"
	"api-rest/models"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetUsers(c fiber.Ctx) error {
	var users []models.User

	result := database.DB.Preload("Tasks").Find(&users)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudieron cargar los datos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": users,
	})
}

func GetUser(c fiber.Ctx) error {
	id := c.Params("id")

	var user models.User

	result := database.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Usuario no encontrado",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	database.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}

func CreateUser(c fiber.Ctx) error {
	var user models.User

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos en el cuerpo de la petición",
		})
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudo crear el usuario: " + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuario creado con éxito",
		"data":    user,
	})
}

func DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	result := database.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Usuario no encontrado",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	result = database.DB.Unscoped().Delete(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudo eliminar el registro" + result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Usuario eliminado",
	})
}
