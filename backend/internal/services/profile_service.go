package services

import (
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type ProfileService struct {
	Repository repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) (ProfileService) {
	return ProfileService{
		Repository: repo,
	}
}

func (service ProfileService) GetName() (string, error) {
	name, err := service.Repository.GetName()
	if err != nil {
		return "", err
	}

	return name, nil
}