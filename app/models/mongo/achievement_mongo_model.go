package mongo

import "time"

type Achievement struct {
	ID            string                 `bson:"_id,omitempty"`
	StudentID     string                 `bson:"studentId"`
	AchievementType string               `bson:"achievementType"`
	Title         string                 `bson:"title"`
	Description   string                 `bson:"description"`
	Details       map[string]interface{} `bson:"details"`
	Attachments   []Attachment           `bson:"attachments"`
	Tags          []string               `bson:"tags"`
	Points        int                    `bson:"points"`
	CreatedAt     time.Time              `bson:"createdAt"`
	UpdatedAt     time.Time              `bson:"updatedAt"`
}

type Attachment struct {
	FileName   string    `bson:"fileName"`
	FileURL    string    `bson:"fileUrl"`
	FileType   string    `bson:"fileType"`
	UploadedAt time.Time `bson:"uploadedAt"`
}
