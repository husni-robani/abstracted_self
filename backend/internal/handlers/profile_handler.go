package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
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

	profileData, err := handler.Service.GetProfileData(req.Name, req.Summary, req.Bio, req.Taglines, req.Resume, req.SkillSet)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get profile name", err)
		return
	}

	c.Header("Content-Type", "application/json")
	if req.Cache {
		c.Header("Cache-Control", "public, max-age=86400")
	}


	response.Success(c, http.StatusOK, "Get profile data success", profileData)
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

func (handler ProfileHandler) AddSkillType(c *gin.Context) {
	var reqBody requests.AddProfileSkillSetType

	// bind request body
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		logger.Error.Printf("Bind request body failed: %#v", err)
		response.Error(c, http.StatusBadRequest, "Bad Request", nil)
		return
	}

	// validate request body data
	invalidFieldErrors, err := utils.ValidateStruct(reqBody)
	if err != nil {
		logger.Error.Printf("Validate request body failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("Invalid body request: %#v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "Invalid Data", invalidFieldErrors)
		return
	}

	// add skill type
	if err := handler.Service.AddSkillSetType(reqBody); err != nil {
		if err == models.ErrSkillTypeDuplicate {
			logger.Info.Println(err)
			response.Error(c, http.StatusConflict, err.Error(), nil)
			return
		}
		logger.Error.Printf("Add profile skill type failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	response.Success(c, http.StatusOK, "add profile skill type successful!", nil)
}

func (handler ProfileHandler) AddSkill(c *gin.Context){
	var reqBody requests.AddProfileSkill
	// get skill_name, is_most_used, file
	if err := c.Bind(&reqBody); err != nil {
		logger.Error.Printf("Bind request body failed: %#v", err)
		response.Error(c, http.StatusBadRequest, "Bad Request", nil)
		return
	}

	// validate data
	invalidFieldErrors, err := utils.ValidateStruct(reqBody)
	if err != nil {
		logger.Error.Printf("Validate request body failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	// get request multipart form file
	reqBody.IconFile, err = c.FormFile("icon_file")
	if err != nil {
		if err == http.ErrMissingFile {
			reqBody.IconFile = nil
		}else {
			logger.Error.Println("failed to get request form file: ", err)
			response.Error(c, http.StatusInternalServerError, "internal server error", nil)
			return
		}
	}
	// validate file
	if reqBody.IconFile != nil {
		if err := utils.ValidateFile(reqBody.IconFile, []string{"image/svg+xml"}, 300 << 10); err != nil {
			invalidFieldErrors["icon_file"] = err.Error()
		}
	}

	// return all validation errors
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("Invalid body request: %#v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid body request", invalidFieldErrors)
		return
	}

	// add skill
	if err := handler.Service.AddSkill(reqBody); err != nil {
		if err == models.ErrSkillDuplicate {
			logger.Info.Println(err.Error())
			response.Error(c, http.StatusConflict, err.Error(), nil)
			return
		}
		if err == models.ErrSkillTypeNotFound{
			logger.Info.Println(err.Error())
			response.Error(c, http.StatusNotFound, err.Error(), nil)
			return
		}

		logger.Error.Printf("Add skill failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "Internal Server Error", nil)
		return
	}

	response.Success(c, http.StatusOK, "Add skill successful", nil)
}

func (handler ProfileHandler) RemoveSkillType(c *gin.Context) {
	var reqBody requests.RemoveProfileSkillSetType
	if err := c.Bind(&reqBody); err != nil {
		logger.Error.Printf("bind request body failed: %#v", err)
		response.Error(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	invalidFieldErrors, err := utils.ValidateStruct(reqBody)
	if err != nil {
		logger.Error.Printf("body request validation failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("Invalid data: %#v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFieldErrors)
		return
	}

	if err := handler.Service.RemoveSkillType(reqBody.TypeName); err != nil {
		if err == models.ErrSkillTypeNotFound {
			logger.Info.Printf("%#v: %s", err, reqBody.TypeName)
			response.Error(c, http.StatusNotFound, "type name not found", nil)
			return
		}

		logger.Error.Printf("Remove skill type failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "remove skill type successful", nil)
}

func (handler ProfileHandler) RemoveSkill(c *gin.Context) {
	var reqBody requests.RemoveProfileSkill

	if err := c.Bind(&reqBody); err != nil {
		logger.Error.Printf("bind request body failed: %#v", err)
		response.Error(c, http.StatusBadRequest, "bad request", nil)
		return
	}

	invalidFieldErrors, err := utils.ValidateStruct(reqBody)
	if err != nil {
		logger.Error.Printf("body request validation failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("Invalid data: %#v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFieldErrors)
		return
	}

	if err := handler.Service.RemoveSkill(reqBody.SkillName); err != nil {
		if err == models.ErrSkillNotFound {
			logger.Info.Printf("%#v: %s", err, reqBody.SkillName)
			response.Error(c, http.StatusNotFound, "skill not found", nil)
			return
		}

		logger.Error.Printf("Remove skill failed: %#v", err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "remove skill successful", nil)
}