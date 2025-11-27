package postgre

import "time"

type Role struct {
	ID          string    `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"size:50;unique;not null"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time
}