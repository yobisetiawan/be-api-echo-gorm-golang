package models

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;index"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"` // Soft delete with gorm.Model convention
}

type BaseNoDateJSONModel struct {
	ID        int64          `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime;index"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime;index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"` // Soft delete with gorm.Model convention
}
