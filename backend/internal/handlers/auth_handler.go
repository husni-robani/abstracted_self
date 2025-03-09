package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return AuthHandler{
		service: authService,
	}
}

func (handler AuthHandler) Login(c *gin.Context) {
	var loginRequest requests.LoginRequest

	err := c.BindJSON(&loginRequest)
	if err != nil {
		logger.Error.Printf("failed bind request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "login failed", nil)
		return
	}

	accessToken, err := handler.service.LoginService(loginRequest.Username, loginRequest.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredetials) {
			response.Error(c, http.StatusUnauthorized, err.Error(), nil)		
		}else {
			response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		}
		return
	}

	data := map[string]string{"token": accessToken}

	response.Success(c, http.StatusOK, "login successful", data)
}

func (handler AuthHandler) Logout(c *gin.Context) {
	err := os.Setenv("AUTH_TOKEN", "")	
	if err != nil {
		logger.Error.Printf("failed to set env: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return 
	}

	response.Success(c, http.StatusOK, "logout successful", nil)
}