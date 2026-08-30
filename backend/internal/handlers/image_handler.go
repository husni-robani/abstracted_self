package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type ImageHandler struct {
	Service services.ImageService
}

func NewImageHandler(imageService services.ImageService) ImageHandler {
	return ImageHandler{
		Service: imageService,
	}
}

func (handler ImageHandler) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid data", map[string]string{"file": "required"})
		return
	}

	if err := utils.ValidateFile(file, []string{"image/jpeg", "image/png"}, 300<<10); err != nil {
		logger.Error.Printf("invalid image file: %v", err)
		response.Error(c, http.StatusBadRequest, "invalid image", err.Error())
		return
	}

	image, err := handler.Service.UploadImage(file)
	if err != nil {
		logger.Error.Printf("error upload image: %v", err.Error())
		response.Error(c, http.StatusInternalServerError, "upload image failed", nil)
		return
	}

	response.Success(c, http.StatusCreated, "image uploaded successfully", image)
}

func (handler ImageHandler) GetImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error.Printf("error converting id: %v", err.Error())
		response.Error(c, http.StatusBadRequest, "invalid id", "invalid image ID")
		return
	}

	image, err := handler.Service.GetImageByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			response.Error(c, http.StatusNotFound, "image not found", nil)
			return
		}
		response.Error(c, http.StatusInternalServerError, "internal server error", nil)
		return
	}

	cwd, _ := os.Getwd()
	fullPath := filepath.Join(cwd, os.Getenv("IMAGES_STORAGE_PATH"), image.FileName)
	if _, err := os.Stat(fullPath); err != nil {
		logger.Error.Printf("image file not found: %v", fullPath)
		response.Error(c, http.StatusNotFound, "File Not Found", nil)
		return
	}

	c.Header("Content-Type", image.MimeType)
	c.Header("Cache-Control", "public, max-age=86400")
	c.File(fullPath)
}
