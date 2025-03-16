package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
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

func (handler ExperienceHandler) CreateExperience(c *gin.Context) {
	// get request body
	createExperienceRequest := requests.CreateExperienceRequest{}
	if err := c.ShouldBind(&createExperienceRequest); err != nil {
		logger.Error.Printf("failed bind request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// validate request body
	invalidErrors, err := utils.ValidateStruct(createExperienceRequest)
	if err != nil {
		logger.Error.Println(err.Error())
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

		// return all invalid errors
	if len(invalidErrors) >= 1 {
		logger.Info.Printf("invalid body request: %v", invalidErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidErrors)
		return
	}

	// insert to database
	if err := handler.service.CreateExperience(createExperienceRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// return success response
	response.Success(c, http.StatusCreated, "create experience successful", nil)
}