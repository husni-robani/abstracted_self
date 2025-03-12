package services

import (
	"errors"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

var ErrTypeNotExist = errors.New("type not exist")

type TechnologyService struct {
	repo repositories.TechnologyRepository
}

func NewTechnologyService(repo repositories.TechnologyRepository) TechnologyService {
	return TechnologyService{
		repo: repo,
	}
}

func (service TechnologyService) GetTechnologies() ([]models.Technology, error) {
	technolgies, err := service.repo.GetAllTechnologies()
	if err != nil {
		return nil, err
	}

return technolgies, err
}

func (service TechnologyService) CreateTechnologies(technologies []requests.CreateTechnologyRequest) error {
	// check is type is exist
	for _, tech := range technologies {
		exist, err := service.repo.IsTypeExist(tech.TypeId)
		if err != nil {
			return err
		}

		if !exist {
			logger.Error.Printf("type id %v is not exist", tech.TypeId)
			return ErrTypeNotExist
		}
	}

	if err := service.repo.CreateTechnologies(technologies); err != nil {
		return err
	}

	return nil
}

func (service TechnologyService) DeleteTechnologyByID(id int) error {
	if err := service.repo.DeleteTechnologyByID(id); err != nil {
		return err
	}

	return nil
}