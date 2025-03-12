package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type TechnologyHandler struct {
	service services.TechnologyService
}

func NewTechnologyHandler(service services.TechnologyService) TechnologyHandler {
	return TechnologyHandler{
		service: service,
	}
}

func (handler TechnologyHandler) GetTechnologies(c *gin.Context) {
	technologies, err := handler.service.GetTechnologies()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	
	response.Success(c, http.StatusOK, "get technologies successful", technologies)
}

func (handler TechnologyHandler) CreateTechnologies(c *gin.Context) {
	var technologies []requests.CreateTechnologyRequest
	err := c.ShouldBindJSON(&technologies)
	if err != nil {
		logger.Error.Printf("error bind request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	if err := handler.service.CreateTechnologies(technologies); err != nil {
		if errors.Is(err, services.ErrTypeNotExist) {
			response.Error(c, http.StatusBadRequest, "type does not exist", err.Error())
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "create technologies successful", nil)
}

func (handler TechnologyHandler) DeleteTechnology(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed convert id: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	if err := handler.service.DeleteTechnologyByID(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "delete technology successful", nil)
}