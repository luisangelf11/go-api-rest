package services

import (
	"api-rest/database"
	"api-rest/models"
	"log"
)

func SaveActivityInBackground(action string, desc string) {
	log.Println("🔄 [Goroutine] Escribiendo auditoría en segundo plano...")
	logEntry := models.ActivityLog{
		Action:      action,
		Description: desc,
	}

	if err := database.DB.Create(&logEntry).Error; err != nil {
		log.Println("❌ [Goroutine] Error de auditoria:", err)
		return
	}

	log.Printf("💾 [Goroutine] Auditoría guardada con ID %d", logEntry.ID)
}
