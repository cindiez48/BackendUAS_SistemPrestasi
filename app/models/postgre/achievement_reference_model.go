package postgre

import "time"

type AchievementReference struct {
	ID                string    `gorm:"type:uuid;primaryKey"`
	StudentID         string    `gorm:"type:uuid"`
	MongoAchievementID string    `gorm:"size:24;not null"`
	Status            string    `gorm:"type:enum('draft','submitted','verified','rejected')"`
	SubmittedAt       *time.Time
	VerifiedAt        *time.Time
	VerifiedBy        *string   `gorm:"type:uuid"`
	RejectionNote     string    `gorm:"type:text"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}