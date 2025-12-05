package postgre

import (
	pmodel "backenduas_sistemprestasi/app/models/postgre"
)

type AchievementReferenceRepository struct{}
type StudentRepository struct{}

func NewAchievementReferenceRepository() *AchievementReferenceRepository {
	return &AchievementReferenceRepository{}
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (r *AchievementReferenceRepository) Create(ref *pmodel.AchievementReference) error {
	// stub: no-op
	return nil
}

func (r *AchievementReferenceRepository) UpdateStatus(refID, status, note, verifier string) error {
	// stub: no-op
	return nil
}

func (r *AchievementReferenceRepository) MarkDeleted(refID string) error {
	// stub: no-op
	return nil
}

func (r *StudentRepository) GetByUserID(userID string) (*pmodel.Student, error) {
	// stub: not found
	return nil, nil
}
