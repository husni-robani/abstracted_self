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
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type BlogHandler struct {
	Service services.BlogService
	Validator validator.Validate
}

func NewBlogHandler(blogService services.BlogService) BlogHandler{
	validator := validator.New(validator.WithRequiredStructEnabled())
	return BlogHandler{
		Service: blogService,
		Validator: *validator,
	}
}

func (handler BlogHandler) GetBlogs(c *gin.Context) {
	blogs, err := handler.Service.GetAllBlogs()
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
		// error not found
		if errors.Is(err, sql.ErrNoRows){
			response.Error(c, http.StatusNotFound, "blog not found", nil)
			return
		}else {
			response.Error(c, http.StatusInternalServerError, "internal server error", nil)
			return
		}
	}

	response.Success(c, http.StatusOK, "get blog successfully", blog)

}

func (handler BlogHandler) CreateBlog(c *gin.Context){	
	// get request body
	var blogData requests.CreateBlogRequest
	if err := c.ShouldBind(&blogData); err != nil {
		logger.Error.Printf("error binding request body: %v", err.Error())
		response.Error(c, http.StatusInternalServerError, "create blog failed", nil)
		return
	}

	invalidFieldErrors, err := utils.ValidateStruct(blogData)
	if err != nil {
		logger.Error.Println(err)
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	// get and validate image
	file, err := c.FormFile("file")
	if err != nil {
		invalidFieldErrors["image"] = "required"
	}else {
		blogData.ImageFile = file
		logger.Info.Printf("file received: %s\n", file.Filename)
		
		if err := utils.ValidateFile(blogData.ImageFile, []string{"image/jpeg", "image/png"}, 300 << 10); err != nil{
			invalidFieldErrors["image"] = err.Error()
		}
	}

	// return all validation errors
	if len(invalidFieldErrors) >= 1 {
		logger.Info.Printf("invalid body request: %v", invalidFieldErrors)
		response.Error(c, http.StatusBadRequest, "invalid data", invalidFieldErrors)
		return
	}

	err = handler.Service.CreateBlog(blogData)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "create blog failed", nil)
		return
	}

	response.Success(c, http.StatusCreated, "blog created!", nil)
}

func (handler BlogHandler) DeleteBlog(c *gin.Context){
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
		}else {
			response.Error(c, http.StatusInternalServerError, "delete blog failed", nil)
			return
		}
	}

	response.Success(c, http.StatusOK, "delete blog successfully", nil)
}

func (blogHandler BlogHandler) UpdateBlog(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("error converting id: %v", err.Error())
		response.Error(c, http.StatusBadRequest, "invalid id", err.Error())
		return
	}

	var req requests.UpdateBlogRequest
	if err := c.ShouldBind(&req); err != nil {
		logger.Error.Printf("error binding request body: %v", err)
		response.Error(c, http.StatusInternalServerError, "update blog failed", nil)
		return
	}

	// validate image
	if req.ImageFile != nil {
		if err := utils.ValidateFile(req.ImageFile, []string{"image/jpeg", "image/png"}, 300 << 10); err != nil{
			logger.Error.Printf("invalid image file: %v", err)
			response.Error(c, http.StatusBadRequest, "invalid image", err.Error())
			return
		}
	}

	updatedBlog, err := blogHandler.Service.UpdateBlog(id, req)
	if err != nil {
		// error not found
		if errors.Is(err, sql.ErrNoRows){
			response.Error(c, http.StatusNotFound, "blog not found", nil)
			return
		}else {
			response.Error(c, http.StatusInternalServerError, "internal server error", nil)
			return
		}
	}
	
	response.Success(c, http.StatusOK, "update blog successful", updatedBlog)
}