package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type ProfileHandler struct {
	Service services.ProfileService
}

func NewProfileHandler(service services.ProfileService) (ProfileHandler) {
	return ProfileHandler{
		Service: service,
	}
}


func (handler ProfileHandler) GetProfileData(c *gin.Context) {
	var req requests.GetProfileRequest
	if err := c.Bind(&req); err != nil {
		logger.Error.Printf("Failed to bind query string: %v", err)
		response.Error(c, http.StatusBadRequest, "query is not accepted", err)
		return
	}

	profileData, err := handler.Service.GetProfileData(req.Name, req.Summary, req.Bio, req.Taglines, req.Resume, req.Skills)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get profile name", err)
		return
	}

	c.Header("Content-Type", "application/json")
	c.Header("Cache-Control", "public, max-age=86400")

	response.Success(c, http.StatusOK, "Get profile name success", profileData)
}