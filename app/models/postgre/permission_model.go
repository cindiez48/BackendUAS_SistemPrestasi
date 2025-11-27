package postgre

import "time"

type Permission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Name        string `gorm:"size:100;unique;not null"`
	Resource    string `gorm:"size:50;not null"`
	Action      string `gorm:"size:50;not null"`
	Description string `gorm:"type:text"`
	CreatedAt   time.Time
}