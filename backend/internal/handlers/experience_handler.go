package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

	isCache := c.Request.URL.Query().Get("cache") == "true"

	c.Header("Content-Type", "application/json")
	if isCache {
		c.Header("Cache-Control", "public, max-age=86400")
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

func (handler ExperienceHandler) UpdateExperience(c *gin.Context) {
	// get id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed to convert parameter id: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// bind request
	var updateExperienceRequest requests.UpdateExperienceRequest
	if err := c.ShouldBind(&updateExperienceRequest); err != nil {
		logger.Error.Printf("failed to bind body request: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// insert to database
	if err := handler.service.UpdateExperience(id, updateExperienceRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	
	// return response success
	response.Success(c, http.StatusOK, "update experience successful", nil)
}

func (handler ExperienceHandler) GetExperienceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed to convert id to int: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	experience, err := handler.service.GetExperienceById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(c, http.StatusNotFound, fmt.Sprintf("experience with id %v not found", id), nil)
			return 
		}

		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "get experience successful", experience)
}

func (handler ExperienceHandler) DeleteExperience(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed to convert id to int: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	if err := handler.service.DeleteExperience(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "delete experience successful", nil)
}