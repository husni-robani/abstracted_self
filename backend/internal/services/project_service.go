package services

import "github.com/husni-robani/abstracted_self/backend/internal/repositories"

type ProjectService struct {
	repo repositories.ProjectRepository
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return ProjectService{repo: repo}
}