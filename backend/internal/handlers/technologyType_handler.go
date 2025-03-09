package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type TechnologyTypeHandler struct {
	service services.TechnologyTypeService
}

func NewTechnologyTypeHandler(service services.TechnologyTypeService) TechnologyTypeHandler {
	return TechnologyTypeHandler{
		service: service,
	}
}

func (handler TechnologyTypeHandler) GetTypes(c *gin.Context) {
	techTypes, err := handler.service.GetAllTypes()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "get technology types successful", techTypes)
}

func (handler TechnologyTypeHandler) CreateTypes(c *gin.Context) {
	var createTypesRequest requests.CreateTypesRequest
	
	err := c.ShouldBindJSON(&createTypesRequest)
	if err != nil {
		logger.Error.Printf("failed binding request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// validate request body
	invalidFields, err := utils.ValidateStruct(createTypesRequest)
	if err != nil {
		logger.Error.Println(err.Error())
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	if len(invalidFields) >= 1 {
		logger.Error.Printf("invalid request data: %v", invalidFields)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFields)
		return
	}

	err = handler.service.CreateTypes(createTypesRequest.TypeNames)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusCreated, "create type successful", nil)
}

func (handler TechnologyTypeHandler) DeleteTypeByID(c *gin.Context) {
	typeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("failed convert param: %v", err)
		response.Error(c, http.StatusBadRequest, "invalid id parameter", nil)
		return
	}

	err = handler.service.DeleteTypeByID(typeId)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "delete technology type successful", nil)
}