package handlers

import "github.com/husni-robani/abstracted_self/backend/internal/services"

type ProjectHandler struct {
	service services.ProjectService
}

func NewProjectHandler(service services.ProjectService) ProjectHandler {
	return ProjectHandler{service: service}
}