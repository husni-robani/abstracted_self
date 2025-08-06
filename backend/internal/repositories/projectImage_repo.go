package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type ProjectImageRepository struct {
	db *sql.DB
}

func NewProjectImageRepository(db *sql.DB) ProjectImageRepository {
	return ProjectImageRepository{db: db}
}

func (repo ProjectImageRepository) AddProjectImages(projectId int, images []multipart.FileHeader) error {
	query := "INSERT INTO project_images (project_id, file_name, file_size, mime_type, image_url) VALUES"

	for i, image := range images {
		image_url := fmt.Sprintf("%s:%s/storage/project_images/%s", os.Getenv("BACKEND_HOST"), os.Getenv("SERVER_PORT"), image.Filename)
		values := fmt.Sprintf("('%d', '%s', '%d', '%s', '%s')", projectId, image.Filename, image.Size, filepath.Ext(image.Filename), image_url)
		
		if i != len(images) - 1 {
			values += ","
		}

		query += values
	}

	result, err := repo.db.Exec(query)
	if err != nil {
		logger.Error.Printf("failed to insert project_image to database: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed to get rows affected: %v", rowsAffected)
		return err
	}

	logger.Info.Printf("create project images successful | rows affected: %d", rowsAffected)

	return nil
}

func (repo ProjectImageRepository) GetImagesByProjectId(projectId int) ([]models.ProjectImage, error) {
	rows, err := repo.db.Query("SELECT id, file_name, file_size FROM project_images WHERE project_id = $1", projectId)
	if err != nil {
		logger.Error.Printf("failed to exec query select: %v", err)
		return nil, err
	}

	var images []models.ProjectImage

	for rows.Next() {
		var image models.ProjectImage

		if err := rows.Scan(&image.Id, &image.FileName, &image.FileSize); err != nil {
			logger.Error.Printf("failed to scan data: %v", err)
			return nil, err
		}

		images = append(images, image)
	}


	return images, nil
}

func (repo ProjectImageRepository) GetImageById(id int) (models.ProjectImage, error) {
	var projectImage models.ProjectImage

	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	
	row := repo.db.QueryRowContext(ctx, "SELECT id, file_name FROM project_images WHERE id = $1;", id)

	if err := row.Scan(&projectImage.Id, &projectImage.FileName); err != nil {
		logger.Error.Printf("scan project_images failed: %v", err)
		return projectImage, err
	}

	return projectImage, nil
}


func (repo ProjectImageRepository) DeleteProjectImageByID(id ...int) error {
	if len(id) < 1 {
		err := errors.New("argument cannot be empty, at least 1 argument needed")
		logger.Error.Printf("Delete project image is failed: %v", err)
		return err
	}

	ids := ""
	for i, v := range id {
		if i == 0 {
			ids += fmt.Sprintf("%v", v)
			continue
		}
		ids += fmt.Sprintf(", %d", v)
	}

	query := fmt.Sprintf("DELETE FROM project_images WHERE id in (%s);", ids)

	timeOut, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := repo.db.ExecContext(timeOut, query)
	if err != nil {
		logger.Error.Printf("delete query execution is failed: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	logger.Info.Printf("delete project image is succeeded | Rows affected: %d", rowsAffected)

	return nil
}