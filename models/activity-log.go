package models

import "gorm.io/gorm"

type ActivityLog struct {
	gorm.Model
	Action      string `json:"action"`
	Description string `json:"description"`
}
