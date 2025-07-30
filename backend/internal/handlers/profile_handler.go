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
	if req.Cache {
		c.Header("Cache-Control", "public, max-age=86400")
	}


	response.Success(c, http.StatusOK, "Get profile name success", profileData)
}

func (handler ProfileHandler) UpdateProfile(c *gin.Context){
	var reqBody requests.UpdateProfileRequest

	// get form data
	if err := c.Bind(&reqBody); err != nil {
		logger.Error.Println("Failed to bind body request: ", err)
		response.Error(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// validate data
	invalidFieldErrors, err := utils.ValidateStruct(reqBody)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
		return 
	}

	// get request multipart form file
	reqBody.ResumeFile, err = c.FormFile("resume_file")
	if err != nil {
		if err == http.ErrMissingFile {
			reqBody.ResumeFile = nil
		}else {
			logger.Error.Println("failed to get request form file: ", err)
			response.Error(c, http.StatusInternalServerError, "internal server error", nil)
			return
		}
	}
	// validate file
	if reqBody.ResumeFile != nil {
		if err := utils.ValidateFile(reqBody.ResumeFile, []string{"application/pdf"}, 300 << 10); err != nil {
			invalidFieldErrors["resume_file"] = err.Error()
		}
	}

	// return all validation errors
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("Invalid body request: %#v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid body request", invalidFieldErrors)
		return
	}

	// update profile
	if err := handler.Service.UpdateProfileData(reqBody); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error(), nil)
	}

	response.Success(c, http.StatusOK, "Update profile successful!", nil)
}