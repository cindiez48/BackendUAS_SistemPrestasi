package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Achievement struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	StudentID      string             `bson:"studentId"`
	AchievementType string            `bson:"achievementType"`
	Title          string             `bson:"title"`
	Description    string             `bson:"description"`

	Details        map[string]interface{} `bson:"details"`

	Attachments []struct {
		FileName  string    `bson:"fileName"`
		FileUrl   string    `bson:"fileUrl"`
		FileType  string    `bson:"fileType"`
		UploadedAt time.Time `bson:"uploadedAt"`
	} `bson:"attachments"`

	Tags      []string  `bson:"tags"`
	Points    float64   `bson:"points"`
	IsDeleted bool      `bson:"isDeleted"`
	DeletedAt *time.Time `bson:"deletedAt"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}