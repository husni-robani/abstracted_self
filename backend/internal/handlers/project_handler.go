package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type ProjectHandler struct {
	service services.ProjectService
}

func NewProjectHandler(service services.ProjectService) ProjectHandler {
	return ProjectHandler{service: service}
}

func (handler ProjectHandler) CreateProject(c *gin.Context) {
	var project_request requests.CreateProjectRequest
	// bind request body and image
	if err := c.ShouldBind(&project_request); err != nil {
		logger.Error.Printf("failed bind request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "create project failed", nil)
		return
	}

	// validate image and request body
	fieldErrors, err := utils.ValidateStruct(project_request)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
		// validate image in more detail
	for _, image := range project_request.Images {
		if err := utils.ValidateFile(&image, []string{"image/jpeg", "image/png"}, 300 << 10); err != nil {
			fieldErrors[fmt.Sprintf("image(%s)", image.Filename)] = err.Error()
		}
	}

	// return all invalid field errors
	if len(fieldErrors) >= 1 {
		logger.Info.Printf("Invalid form request: %v", fieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", fieldErrors)
		return
	}
 
	// send request and image to service
	if err := handler.service.CreateNewProject(project_request); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusCreated, "create project successful", nil)
}

func (handler ProjectHandler) GetProjects(c *gin.Context) {
	projects, err := handler.service.GetAllProjectsWithImages()

	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "get projects successful", projects)
}