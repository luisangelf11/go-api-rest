package handlers

import (
	"api-rest/database"
	"api-rest/models"
	"api-rest/services"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func GetTasks(c fiber.Ctx) error {
	var tasks []models.Task

	result := database.DB.Find(&tasks)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se pudieron cargar los datos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tasks,
	})
}

func GetTask(c fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task

	result := database.DB.First(&task, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "La tarea no fue encontrada",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": task,
	})
}

func CreateTask(c fiber.Ctx) error {
	var task models.Task

	if err := c.Bind().Body(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos en el cuerpo de la petición",
		})
	}

	result := database.DB.Create(&task)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se ha podido crear la tarea",
		})
	}

	go services.SaveActivityInBackground("TASK_CREATED", "Se creó la tarea: "+task.Title)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tarea creada con exito",
		"data":    task,
	})
}

func DeleteTask(c fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task

	result := database.DB.First(&task, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "La tarea no fue encontrada",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	result = database.DB.Unscoped().Delete(&task)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se ha podido eliminar la tarea",
		})
	}

	go services.SaveActivityInBackground("TASK_DELETED", "Se eliminó la tarea con id: "+id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "La tarea fue eliminada",
	})
}

func UpdateTask(c fiber.Ctx) error {
	id := c.Params("id")
	var task models.Task
	var body models.Task

	result := database.DB.First(&task, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "La tarea no fue encontrada",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error interno del servidor",
		})
	}

	if err := c.Bind().Body(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Datos inválidos en el cuerpo de la petición",
		})
	}

	result = database.DB.Model(&task).Updates(&body)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No se ha podido actualizar la tarea",
		})
	}

	go services.SaveActivityInBackground("TASK_UPDATED", "Se actualizó la tarea con id: "+id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "La tarea fue actualizada",
		"data":    task,
	})
}
