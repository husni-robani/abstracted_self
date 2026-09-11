package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	identifier, err := c.Cookie("visitor_identifier")
	if err != nil || identifier == "" {
		response.Error(c, http.StatusBadRequest, "visitor_identifier cookie is required", nil)
		return
	}

	if _, err := uuid.Parse(identifier); err != nil {
		logger.Info.Printf("Invalid visitor identifier: %v", err)
		response.Error(c, http.StatusBadRequest, "invalid visitor_identifier", nil)
		return
	}

	created, err := handler.service.ProfileVisitor(identifier, c)
	if err != nil {
		logger.Error.Printf("Failed to record profile visit: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	if !created {
		response.Success(c, http.StatusOK, "visitor already recorded", nil)
		return
	}

	response.Success(c, http.StatusCreated, "Success", nil)
}

func (handler DailyVisitHandler) GetDailyVisitCounts(c *gin.Context) {
	var query requests.DailyVisitsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.Error.Printf("Failed to bind query string: %v", err)
		response.Error(c, http.StatusBadRequest, "query is not accepted", err)
		return
	}

	invalidFields, err := utils.ValidateStruct(query)
	if err != nil {
		logger.Error.Printf("Validation failed: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	if len(invalidFields) >= 1 {
		logger.Info.Printf("Invalid field: %#v", invalidFields)
		response.Error(c, http.StatusBadRequest, "Invalid Data", invalidFields)
		return
	}

	if err := utils.ValidateDateRange(query.StartDate, query.EndDate); err != nil {
		logger.Info.Printf("Invalid date range: %v", err)
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	counts, err := handler.service.GetDailyVisitCounts(query.StartDate, query.EndDate)
	if err != nil {
		logger.Error.Printf("Failed to get daily visit counts: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	response.Success(c, http.StatusOK, "Get daily visit counts successful", counts)
}

func (handler DailyVisitHandler) GetDailyVisits(c *gin.Context) {
	var query requests.DailyVisitsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		logger.Error.Printf("Failed to bind query string: %v", err)
		response.Error(c, http.StatusBadRequest, "query is not accepted", err)
		return
	}

	invalidFields, err := utils.ValidateStruct(query)
	if err != nil {
		logger.Error.Printf("Validation failed: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}
	if len(invalidFields) >= 1 {
		logger.Info.Printf("Invalid field: %#v", invalidFields)
		response.Error(c, http.StatusBadRequest, "Invalid Data", invalidFields)
		return
	}

	if err := utils.ValidateDateRange(query.StartDate, query.EndDate); err != nil {
		logger.Info.Printf("Invalid date range: %v", err)
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	page, limit := utils.NormalizePagination(query.Page, query.Limit)

	visits, pagination, err := handler.service.GetDailyVisits(query.StartDate, query.EndDate, page, limit)
	if err != nil {
		logger.Error.Printf("Failed to get daily visits: %v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	data := gin.H{
		"daily_visits": visits,
		"pagination":   pagination,
	}

	response.Success(c, http.StatusOK, "Get daily visits successful", data)
}