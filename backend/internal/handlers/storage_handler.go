package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/response"
)

type StorageHandler struct {

}

func NewStorageHandler() (StorageHandler) {
	return StorageHandler{}
}

func (StorageHandler) GetProjectImage(c *gin.Context) {
	cwd, _ := os.Getwd()
	file_name := c.Param("file_name")
	full_path := filepath.Join(cwd, "storage/project_images", file_name)

	// Get file info
	_, err := os.Stat(full_path)
	if err != nil || os.IsNotExist(err){
		response.Error(c, http.StatusNotFound, "File Not Found", fmt.Sprintf("%v is not found", full_path))
		return
	}

	c.File(full_path)
}