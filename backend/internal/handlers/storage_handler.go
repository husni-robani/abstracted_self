package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
)

type StorageHandler struct {

}

func NewStorageHandler() (StorageHandler) {
	return StorageHandler{}
}

func (StorageHandler) GetImages(c *gin.Context) {
	cwd, _ := os.Getwd()
	file_name := c.Param("file_name")
	full_path := filepath.Join(cwd, "storage/images", file_name)

	// Get file info
	_, err := os.Stat(full_path)
	if err != nil || os.IsNotExist(err){
		logger.Error.Printf("File not found error: %v", full_path)
		response.Error(c, http.StatusNotFound, "File Not Found", nil)
		return
	}

	c.Header("Content-Type", "image/jpeg")
	c.Header("Cache-Control", "public, max-age=86400")
	c.File(full_path)
}

func (StorageHandler) GetDocuments(c *gin.Context) {
	cwd, _ := os.Getwd()
	file_name := c.Param("file_name")
	full_path := filepath.Join(cwd, "storage/documents", file_name)	

	// Get file info
	_, err := os.Stat(full_path)
	if err != nil || os.IsNotExist(err){
		response.Error(c, http.StatusNotFound, "File Not Found", fmt.Sprintf("%v is not found", full_path))
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.File(full_path)
}