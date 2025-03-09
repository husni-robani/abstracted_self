package services

import "github.com/husni-robani/abstracted_self/backend/internal/repositories"

type TechnologyTypeService struct {
	repository repositories.TechnologyTypeRepository
}

func NewTechnologyTypeService(repository repositories.TechnologyTypeRepository) TechnologyTypeService {
	return TechnologyTypeService{repository: repository}
}