package postgre

import "time"

type Lecturer struct {
	ID         string    `gorm:"type:uuid;primaryKey"`
	UserID     string    `gorm:"type:uuid"`
	LecturerID string    `gorm:"size:20;unique;not null"`
	Department string    `gorm:"size:100"`
	CreatedAt  time.Time
}