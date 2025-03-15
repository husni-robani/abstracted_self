package handlers

import "github.com/husni-robani/abstracted_self/backend/internal/services"

type ExperienceHandler struct {
	service services.ExperienceService
}

func NewExperienceHandler(service services.ExperienceService) ExperienceHandler {
	return ExperienceHandler{
		service: service,
	}
}