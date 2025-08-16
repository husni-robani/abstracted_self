package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"reflect"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
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

func RemoveFile(dirPath string, fileName string) error {
	err := os.Remove(dirPath + fileName)
	if err != nil {
		logger.Error.Printf("remove file failed: %v", err)
		return err
	}

	logger.Info.Printf("file removed successfully: %s/%s", dirPath, fileName)
	return nil
}

func ValidateFile(fileHeader *multipart.FileHeader, validTypes []string, maxSize int64) error {
	if fileHeader.Size > maxSize {
		return fmt.Errorf("image larger than %d KB",(maxSize / 1024))
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

func GenerateSingleUpdateQuery(tableName string, values map[string]any, condition string) string {
	query := fmt.Sprintf("UPDATE %s SET",tableName)

	i := 0
	for column, value := range values{
		i++
		if i == len(values) {
			if value != nil {
				query += fmt.Sprintf(" %s = '%s' ", column, value)	
			}else {
				query += fmt.Sprintf(" %s = NULL ", column)
			}
			break
		}

		if value != nil {
			query += fmt.Sprintf(" %s = '%s',", column, value)
		}else {
			query += fmt.Sprintf(" %s = NULL,", column, )
		}
	}

	query += condition

	return query
}

func GetToken(bearerToken string) (string, error) {
	bearerSplitted := strings.Split(bearerToken, " ")

	if len(bearerSplitted) != 2 {
		return "", fmt.Errorf("invalid token")
	}

	return bearerSplitted[1], nil
}

func ValidateStruct(structInstance any) (map[string]string, error) {
	argumentType := reflect.ValueOf(structInstance)

	if argumentType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("failed validate struct: invalid argument (expected struct)")
	}

	// validation process
	invalidFields :=  map[string]string{}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(structInstance)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs){
			for _, e := range validateErrs {
				invalidFields[e.Field()] = e.ActualTag()
			}
		}else {
			return nil, err
		}
	}

	return invalidFields, nil
}