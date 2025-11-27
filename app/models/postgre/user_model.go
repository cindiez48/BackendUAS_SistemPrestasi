package postgre

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string         `gorm:"type:uuid;primaryKey"`
	Username     string         `gorm:"size:50;unique;not null"`
	Email        string         `gorm:"size:100;unique;not null"`
	PasswordHash string         `gorm:"size:255;not null"`
	FullName     string         `gorm:"size:100;not null"`
	RoleID       string         `gorm:"type:uuid"`
	IsActive     bool           `gorm:"default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
