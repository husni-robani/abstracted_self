package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type ExperienceHandler struct {
	service services.ExperienceService
}

func NewExperienceHandler(service services.ExperienceService) ExperienceHandler {
	return ExperienceHandler{
		service: service,
	}
}

func (handler ExperienceHandler) GetExperiences(c *gin.Context) {
	experiences, err := handler.service.GetAllExperiences()
	
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed get all experiences", nil)
		return
	}

	response.Success(c, http.StatusOK, "get all experiences successful", experiences)
}