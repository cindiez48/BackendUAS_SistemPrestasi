package mongo

import (
	"context"

	modelmongo "backenduas_sistemprestasi/app/models/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementMongoRepository struct{}

func NewAchievementMongoRepository() *AchievementMongoRepository {
	return &AchievementMongoRepository{}
}

func (r *AchievementMongoRepository) Create(ctx context.Context, ach *modelmongo.Achievement) (primitive.ObjectID, error) {
	// Minimal stub: return a new ObjectID and nil error
	return primitive.NewObjectID(), nil
}

func (r *AchievementMongoRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*modelmongo.Achievement, error) {
	// Return empty achievement as placeholder
	return &modelmongo.Achievement{}, nil
}

func (r *AchievementMongoRepository) SoftDelete(ctx context.Context, id primitive.ObjectID) error {
	// Stub: no-op
	return nil
}
