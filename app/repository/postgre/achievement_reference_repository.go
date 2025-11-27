package postgre

import (
    "database/sql"
    "backenduas_sistemprestasi/database"
    model "backenduas_sistemprestasi/app/models/postgre"
)

type AchievementReferenceRepository struct {
    DB *sql.DB
}

func NewAchievementReferenceRepository() *AchievementReferenceRepository {
    return &AchievementReferenceRepository{
        DB: database.PostgresDB,
    }
}

func (r *AchievementReferenceRepository) Insert(ref *model.AchievementReference) error {
    query := `
        INSERT INTO achievement_references 
        (id, student_id, mongo_achievement_id, status, submitted_at, verified_at,
         verified_by, rejection_note, created_at, updated_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
    `
    _, err := r.DB.Exec(
        query,
        ref.ID, ref.StudentID, ref.MongoAchievementID,
        ref.Status, ref.SubmittedAt, ref.VerifiedAt,
        ref.VerifiedBy, ref.RejectionNote,
        ref.CreatedAt, ref.UpdatedAt,
    )
    return err
}