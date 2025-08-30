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

type DailyVisitHandler struct {
	service services.DailyVisitService
}

func NewDailyVisitHandler(service services.DailyVisitService) DailyVisitHandler {
	return DailyVisitHandler{
		service: service,
	}
}

func (handler DailyVisitHandler) ProfileVisitor(c *gin.Context) {
	var visitRequest requests.VisitRequest

	if err := c.ShouldBind(&visitRequest); err != nil {
		logger.Error.Printf("Failed to bind request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "failed to bind request", nil)
		return
	}
	
	invalidFields, err := utils.ValidateStruct(visitRequest)
	if err != nil {
		logger.Error.Printf("Request validation failed: %v", err)
		response.Error(c, http.StatusInternalServerError, "Validation failed", nil)
		return
	}

	if len(invalidFields) >= 1 {
		logger.Info.Printf("Invalid field: %#v", invalidFields)
		response.Error(c, http.StatusBadRequest, "Invalid Request", invalidFields)
		return
	}


	if err = handler.service.ProfileVisitor(visitRequest); err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	response.Success(c, http.StatusCreated, "Success", nil)
}