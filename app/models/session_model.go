package models

import "time"

type Session struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	AdminID      int64      `json:"admin_id" gorm:"default:null;index"`
	Admin        Admin      `json:"admin,omitempty" gorm:"foreignKey:AdminID;references:ID"` // Define relationship
	UserID       int64      `json:"user_id" gorm:"default:null;index"`
	User         User       `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID"` // Define relationship
	IsActive     bool       `json:"is_active" gorm:"default:false"`
	ExpiredAt    *time.Time `json:"expired_at" gorm:""`
	RefreshToken string     `json:"refresh_token" gorm:"type:varchar(500);default:null;index"`
}
