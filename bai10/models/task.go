package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model 

	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'pending'"` // 'pending', 'completed'
}