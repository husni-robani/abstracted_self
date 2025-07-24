package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func (handler ProfileHandler) GetName(ctx *gin.Context) {
	name, err := handler.Service.GetName()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get profile name", err)
		return
	}

	response.Success(ctx, http.StatusOK, "Get profile name success", name)
}