package services

import (
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type TechnologyTypeService struct {
	repository repositories.TechnologyTypeRepository
}

func NewTechnologyTypeService(repository repositories.TechnologyTypeRepository) TechnologyTypeService {
	return TechnologyTypeService{repository: repository}
}

func (service TechnologyTypeService) GetAllTypes() ([]models.TechnologyType, error) {
	techTypes, err := service.repository.GetAllType()
	if err != nil {
		return nil, err
	}

	return techTypes, nil
}

func (service TechnologyTypeService) CreateTypes(typeNames []string) error {
	err := service.repository.CreateTypes(typeNames)
	if err != nil {
		return err
	}

	return nil
}

func (service TechnologyTypeService) DeleteTypeByID(id int) error {
	err := service.repository.DeleteType(id)
	if err != nil {
		return err
	}

	return nil
}

func (service TechnologyTypeService) GetAllTypesWithTechnologies() ([]models.TechnologyType, error) {
	techTypeWithTechnologies, err := service.repository.GetTypesWithTechnologies()
	if err != nil {
		return nil, err
	}

	return techTypeWithTechnologies, nil
}