package services

import (
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type ImageService struct {
	Repository repositories.ImageRepository
}

func NewImageService(imageRepo repositories.ImageRepository) ImageService {
	return ImageService{Repository: imageRepo}
}

func (service ImageService) UploadImage(fileHeader *multipart.FileHeader) (models.Image, error) {
	extension := filepath.Ext(fileHeader.Filename)
	newFilename := uuid.New().String() + extension

	fileHeader.Filename = newFilename
	mimeType := fileHeader.Header.Get("Content-Type")

	if err := utils.SaveFile(fileHeader, "."+os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
		return models.Image{}, err
	}

	id, err := service.Repository.CreateImage(newFilename, fileHeader.Size, mimeType)
	if err != nil {
		return models.Image{}, err
	}

	return models.Image{
		Id:       id,
		URL:      models.ImageURL(id),
		FileName: newFilename,
		FileSize: int(fileHeader.Size),
		MimeType: mimeType,
	}, nil
}

func (service ImageService) GetImageByID(id int) (models.Image, error) {
	image, err := service.Repository.GetImageByID(id)
	if err != nil {
		return models.Image{}, err
	}

	return image, nil
}

func (service ImageService) DeleteImage(image models.Image) error {
	if err := service.Repository.DeleteImage(image.Id); err != nil {
		return err
	}

	if err := utils.RemoveFile("."+os.Getenv("IMAGES_STORAGE_PATH")+"/", image.FileName); err != nil {
		logger.Error.Printf("error delete image file from storage: %v", err.Error())
		return err
	}

	return nil
}
