package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"slices"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
)

func SaveFile(fileHeader *multipart.FileHeader, saveDir string) error {
	file, err := fileHeader.Open()
	if err != nil {
		logger.Error.Printf("error open file: %v", err.Error())
		return err
	}
	defer file.Close()

	// create file in the specified direcory
	outFile, err := os.Create(fmt.Sprintf("%s/%s", saveDir, fileHeader.Filename))
	if err != nil {
		logger.Error.Printf("error create file: %s", err.Error())
		return err
	}
	defer outFile.Close()

	// copy file content to the new file 
	_, err = io.Copy(outFile, file)
	if err != nil {
		logger.Error.Printf("error copy content to the new file: %s", err.Error())
		return err
	}

	logger.Info.Printf("file saved successfully: %s", outFile.Name())
	return nil
}

func ValidateFile(fileHeader *multipart.FileHeader, validTypes []string, maxSize int64) error {
	if fileHeader.Size > maxSize {
		return errors.New("image larger than 300 KB")
	}

	if  !slices.Contains(validTypes, fileHeader.Header.Get("Content-Type")) {
		var validTypesStr string
		for i, v := range validTypes{
			if i == len(validTypes) - 1{
				validTypesStr += fmt.Sprintf("and '%s'", v)
			}else {
				validTypesStr += fmt.Sprintf("'%s', ", v)
			}
		}
		return fmt.Errorf("content type is invalid, only receives %v", validTypesStr)
	}

	return nil
}