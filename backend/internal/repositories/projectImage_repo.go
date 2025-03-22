package repositories

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"path/filepath"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
)

type ProjectImageRepository struct {
	db *sql.DB
}

func NewProjectImageRepository(db *sql.DB) ProjectImageRepository {
	return ProjectImageRepository{db: db}
}

func (repo ProjectImageRepository) AddProjectImages(projectId int, images []multipart.FileHeader) error {
	query := "INSERT INTO project_images (project_id, file_name, file_size, mime_type) VALUES"

	for i, image := range images {
		values := fmt.Sprintf("('%d', '%s', '%d', '%s')", projectId, image.Filename, image.Size, filepath.Ext(image.Filename))
		
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