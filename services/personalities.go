package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ln0rd/api-golang-persons/errs"
	"github.com/ln0rd/api-golang-persons/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PersonalityService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewPersonalityService(db *gorm.DB, logger *zap.Logger) *PersonalityService {
	return &PersonalityService{db: db, logger: logger}
}

func (s *PersonalityService) CreatePersonality(personality models.Personality) (models.Personality, error) {
	s.logger.Info("Creating personality", zap.Any("personality", personality))

	if err := personality.Validate(); err != nil {
		s.logger.Error("Invalid personality data", zap.Error(err))
		return personality, err
	}

	result := s.db.Create(&personality)
	if result.Error != nil {
		s.logger.Error("Error creating personality", zap.Error(result.Error))
		return personality, errs.NewCustomError("Create Personality Error", http.StatusInternalServerError, result.Error.Error())
	}

	return personality, nil
}

func (s *PersonalityService) UpdatePersonality(id string, personality models.Personality) (models.Personality, error) {
	s.logger.Info("Updating personality", zap.String("id", id))

	if id == "" {
		return personality, errs.NewBadParamError("id is required")
	}

	if err := personality.Validate(); err != nil {
		s.logger.Error("Invalid personality data", zap.Error(err))
		return personality, err
	}

	result := s.db.Save(&personality)
	if result.Error != nil {
		s.logger.Error("Error updating personality", zap.Error(result.Error))
		return personality, errs.NewCustomError("UpdatePersonalityError", http.StatusInternalServerError, result.Error.Error())
	}

	return personality, nil
}

func (s *PersonalityService) DeletePersonality(id string) error {
	if id == "" {
		return errs.NewBadParamError("id is required")
	}

	s.logger.Info("Deleting personality", zap.String("id", id))

	result := s.db.Delete(&models.Personality{}, id)
	if result.Error != nil {
		s.logger.Error("Error deleting personality", zap.Error(result.Error))
		return errs.NewCustomError("DeletePersonalityError", http.StatusInternalServerError, result.Error.Error())
	}

	return nil
}

func (s *PersonalityService) GetAllPersonalities() ([]models.Personality, error) {
	s.logger.Info("Getting all personalities")
	var personalities []models.Personality

	result := s.db.Find(&personalities)
	if result.Error != nil {
		s.logger.Error("Error fetching personalities",
			zap.Error(result.Error))
		return nil, errs.NewCustomError("Internal Server Error", http.StatusInternalServerError, "error fetching personalities")
	}

	s.logger.Debug("Personalities found",
		zap.Int("count", len(personalities)),
		zap.Int64("rows_affected", result.RowsAffected))

	return personalities, nil
}

func (s *PersonalityService) GetPersonalityByID(ID string) (models.Personality, error) {
	s.logger.Info("Getting personality by ID", zap.String("ID", ID))
	var personality models.Personality

	err := s.db.First(&personality, ID).Error
	if err != nil {
		// Check if the error is a record not found error by GORM
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return personality, errs.NewNotFoundError(fmt.Sprintf("personality with ID %s not found", ID))
		}
		s.logger.Error("Error getting personality", zap.Error(err))
		return personality, errs.NewCustomError("Internal Server Error", http.StatusInternalServerError, "error fetching personality")
	}

	s.logger.Debug("Personality found", zap.String("ID", ID))
	return personality, nil
}
