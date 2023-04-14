package models

import "time"

// GormModel is the base model for all models

type GormModel struct {
	ID 	  uint `gorm:"primary_key"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}