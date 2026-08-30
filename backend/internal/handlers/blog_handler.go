package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type BlogHandler struct {
	Service   services.BlogService
	Validator validator.Validate
}

func NewBlogHandler(blogService services.BlogService) BlogHandler {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return BlogHandler{
		Service:   blogService,
		Validator: *validator,
	}
}

func (handler BlogHandler) GetBlogs(c *gin.Context) {
	var published *bool

	switch c.Query("published") {
	case "true":
		value := true
		published = &value
	case "false":
		value := false
		published = &value
	}

	blogs, err := handler.Service.GetAllBlogs(published)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "get blogs failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "get blogs successfully", blogs)
}

func (handler BlogHandler) GetBlogByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("converting error: %#v", err.Error())
		response.Error(c, http.StatusBadRequest, "get blog failed", "invalid blog ID")
		return
	}

	blog, err := handler.Service.GetBlogByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(c, http.StatusNotFound, "blog not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "get blog successfully", blog)
}

func (handler BlogHandler) CreateBlog(c *gin.Context) {
	var blogData requests.CreateBlogRequest
	if err := c.ShouldBind(&blogData); err != nil {
		logger.Error.Printf("error binding request body: %v", err.Error())
		response.Error(c, http.StatusBadRequest, "create blog failed", nil)
		return
	}

	invalidFieldErrors, err := utils.ValidateStruct(blogData)
	if err != nil {
		logger.Error.Println(err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// cover image is required
	file, err := c.FormFile("file")
	if err != nil {
		invalidFieldErrors["file"] = "required"
	} else {
		blogData.ImageFile = file
		logger.Info.Printf("file received: %s\n", file.Filename)

		if err := utils.ValidateFile(blogData.ImageFile, []string{"image/jpeg", "image/png"}, 300<<10); err != nil {
			invalidFieldErrors["file"] = err.Error()
		}
	}

	// return all validation errors
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("invalid body request: %v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFieldErrors)
		return
	}

	if err := handler.Service.CreateBlog(blogData); err != nil {
		if errors.Is(err, models.ErrImageNotFound) {
			response.Error(c, http.StatusBadRequest, "invalid data", map[string]string{"content_image_ids": "one or more images not found"})
			return
		}
		response.Error(c, http.StatusInternalServerError, "create blog failed", nil)
		return
	}

	response.Success(c, http.StatusCreated, "blog created!", nil)
}

func (handler BlogHandler) DeleteBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("error converting id: %v", err)
		response.Error(c, http.StatusBadRequest, "invalid id", err)
		return
	}

	if err := handler.Service.DeleteBlog(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(c, http.StatusNotFound, "blog not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "delete blog failed", nil)
		return
	}

	response.Success(c, http.StatusOK, "delete blog successfully", nil)
}

func (blogHandler BlogHandler) UpdateBlog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("error converting id: %v", err.Error())
		response.Error(c, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	var req requests.UpdateBlogRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error.Printf("error binding request body: %v", err)
		response.Error(c, http.StatusBadRequest, "update blog failed", nil)
		return
	}

	invalidFieldErrors, err := utils.ValidateStruct(req)
	if err != nil {
		logger.Error.Println(err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// get and validate optional replacement cover image
	if file, err := c.FormFile("file"); err == nil {
		req.ImageFile = file
		if err := utils.ValidateFile(req.ImageFile, []string{"image/jpeg", "image/png"}, 300<<10); err != nil {
			invalidFieldErrors["file"] = err.Error()
		}
	} else if err != http.ErrMissingFile {
		logger.Error.Printf("error get image file: %v", err.Error())
		response.Error(c, http.StatusBadRequest, "update blog failed", "invalid image file")
		return
	}

	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("invalid body request: %v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFieldErrors)
		return
	}

	updatedBlog, err := blogHandler.Service.UpdateBlog(id, req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(c, http.StatusNotFound, "blog not found", nil)
			return
		}
		if errors.Is(err, models.ErrImageNotFound) {
			response.Error(c, http.StatusBadRequest, "invalid data", map[string]string{"content_image_ids": "one or more images not found"})
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	response.Success(c, http.StatusOK, "update blog successful", updatedBlog)
}
