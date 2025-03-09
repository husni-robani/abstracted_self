package handlers

import "github.com/husni-robani/abstracted_self/backend/internal/services"

type TechnologyTypeHandler struct {
	service services.TechnologyTypeService
}

func NewTechnologyTypeHandler(service services.TechnologyTypeService) TechnologyTypeHandler {
	return TechnologyTypeHandler{service: service}
}