package service

import (
	"context"
	"time"

	modelmongo "backenduas_sistemprestasi/app/models/mongo"
	pmodel "backenduas_sistemprestasi/app/models/postgre"
	mrepo "backenduas_sistemprestasi/app/repository/mongo"
	postgre "backenduas_sistemprestasi/app/repository/postgre"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AchievementService struct {
	MongoRepo   *mrepo.AchievementMongoRepository
	RefRepo     *postgre.AchievementReferenceRepository
	StudentRepo *postgre.StudentRepository
}

func NewAchievementService() *AchievementService {
	return &AchievementService{
		MongoRepo:   mrepo.NewAchievementMongoRepository(),
		RefRepo:     postgre.NewAchievementReferenceRepository(),
		StudentRepo: postgre.NewStudentRepository(),
	}
}

func (s *AchievementService) CreateAchievement(ctx context.Context, studentID string, ach *modelmongo.Achievement) (refID string, mongoHex string, err error) {
	_, serr := s.StudentRepo.GetByUserID(studentID)
	if serr != nil {
	}

	objID, err := s.MongoRepo.Create(ctx, ach)
	if err != nil {
		return "", "", err
	}

	ref := &pmodel.AchievementReference{
		ID:                 uuid.New().String(),
		StudentID:          studentID,
		MongoAchievementID: objID.Hex(),
		Status:             "draft",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	if err := s.RefRepo.Create(ref); err != nil {
		return "", "", err
	}

	return ref.ID, objID.Hex(), nil
}

func (s *AchievementService) SubmitForVerification(refID string) error {
	err := s.RefRepo.UpdateStatus(refID, "submitted", "", "")
	return err
}

func (s *AchievementService) Verify(refID string, verifierUserID string) error {
	// set verified
	err := s.RefRepo.UpdateStatus(refID, "verified", "", verifierUserID)
	return err
}

func (s *AchievementService) Reject(refID string, note string, verifierUserID string) error {
	err := s.RefRepo.UpdateStatus(refID, "rejected", note, verifierUserID)
	return err
}

func (s *AchievementService) DeleteDraft(ctx context.Context, refID string, mongoHex string) error {
	if mongoHex != "" {
		objID, err := primitive.ObjectIDFromHex(mongoHex)
		if err == nil {
			_ = s.MongoRepo.SoftDelete(ctx, objID)
		}
	}

	if err := s.RefRepo.MarkDeleted(refID); err != nil {
		return err
	}

	return nil
}

func (s *AchievementService) GetAchievementDetail(ctx context.Context, ref pmodel.AchievementReference) (map[string]interface{}, error) {
	out := map[string]interface{}{
		"reference":   ref,
		"achievement": nil,
	}

	if ref.MongoAchievementID == "" {
		return out, nil
	}

	objID, err := primitive.ObjectIDFromHex(ref.MongoAchievementID)
	if err != nil {
		return out, nil
	}

	ach, err := s.MongoRepo.FindByID(ctx, objID)
	if err != nil {
		return out, err
	}

	out["achievement"] = ach
	return out, nil
}
