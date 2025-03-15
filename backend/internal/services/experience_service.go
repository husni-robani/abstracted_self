package services

import "github.com/husni-robani/abstracted_self/backend/internal/repositories"

type ExperienceService struct {
	repo repositories.ExperienceRepo
}

func NewExperienceService(repo repositories.ExperienceRepo) ExperienceService {
	return ExperienceService{
		repo: repo,
	}
}