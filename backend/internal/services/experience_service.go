package services

import (
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type ExperienceService struct {
	repo repositories.ExperienceRepo
}

func NewExperienceService(repo repositories.ExperienceRepo) ExperienceService {
	return ExperienceService{
		repo: repo,
	}
}

func (service ExperienceService) GetAllExperiences() ([]models.Experience, error) {
	experiences, err := service.repo.GetAllExperiences()
	if err != nil {
		return nil, err
	}

	return experiences, nil 
}