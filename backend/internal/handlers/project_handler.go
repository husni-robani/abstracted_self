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

	isCache := c.Request.URL.Query().Get("cache") == "true"

	c.Header("Content-Type", "application/json")
	if isCache {
		c.Header("Cache-Control", "public, max-age=86400")
	}

	response.Success(c, http.StatusOK, "get projects successful", projects)
}

func (handler ProjectHandler) GetProjectById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed to convert id: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	project, err := handler.service.GetProjectById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Success(c, http.StatusNotFound, "project not found", nil)
			return
		}

		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "get project successful", project)
}

func (handler ProjectHandler) DeleteProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("invalid id: %v", err)
		response.Error(c, http.StatusBadRequest, "id invalid!", nil)
		return
	}

	if err := handler.service.DeleteProjectById(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "delete project successful", nil)
}

func (handler ProjectHandler) UpdateProject(c *gin.Context) {
	projectId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("ID convertion failed: %v", err)
		response.Error(c, http.StatusBadRequest, "ID convertion is failed", nil)
		return
	}

	var projectRequest requests.UpdateProjectRequest
	if err := c.Bind(&projectRequest); err != nil {
		logger.Error.Printf("Bind request form failed: %v", err)
		response.Error(c, http.StatusBadRequest, "Bind request form are failed", nil)
		return
	}

	if err := handler.service.UpdateProject(projectId, projectRequest); err != nil {
		logger.Error.Printf("update project is failed: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	response.Success(c, http.StatusOK, "Update project is succeeded", nil)
}