package mongo

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"

    db "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/mongo"
)

type AchievementRepository struct {
    DB *mongo.Database
}

func NewAchievementRepository() *AchievementRepository {
    return &AchievementRepository{
        DB: db.MongoDB,
    }
}

func (r *AchievementRepository) Insert(ctx context.Context, achievement *model.Achievement) error {
    collection := r.DB.Collection("achievements")

    achievement.CreatedAt = time.Now()
    achievement.UpdatedAt = time.Now()

    _, err := collection.InsertOne(ctx, achievement)
    return err
}

func (r *AchievementRepository) FindByID(ctx context.Context, id string) (*model.Achievement, error) {
    collection := r.DB.Collection("achievements")

    var result model.Achievement

    err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
    if err != nil {
        return nil, err
    }

    return &result, nil
}