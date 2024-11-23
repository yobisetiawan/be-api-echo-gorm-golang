package models

import (
	"time"
)

// User represents the user model with specific fields
type User struct {
	BaseModel
	Name             string     `json:"name" gorm:"type:varchar(255);not null"`
	Email            string     `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	Password         string     `json:"-" gorm:"type:varchar(255);not null;index"`
	EmailVerifiedAt  *time.Time `json:"email_verified_at" gorm:""`
	Status           string     `json:"status" gorm:"type:varchar(255);default:'active';index"`
	Role             string     `json:"role" gorm:"type:varchar(255);default:'customer';index"`
	AvatarURL        string     `json:"avatar_url" gorm:"type:varchar(255)"`
	MarkForDeletedAt *time.Time `json:"mark_for_deleted_at" gorm:""`
	NameOld          string     `json:"-" gorm:"type:varchar(255)"`
	EmailOld         string     `json:"-" gorm:"type:varchar(255)"`
}
