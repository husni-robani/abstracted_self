package services

import (
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type ProjectService struct {
	projectRepo repositories.ProjectRepository
	projectImageRepo repositories.ProjectImageRepository
}

func NewProjectService(projectRepo repositories.ProjectRepository, projectImageRepo repositories.ProjectImageRepository) ProjectService {
	return ProjectService{
		projectRepo: projectRepo,
		projectImageRepo: projectImageRepo,
	}
}
