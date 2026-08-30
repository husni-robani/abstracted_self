package repositories

import (
	"database/sql"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/lib/pq"
)

type ImageRepository struct {
	db *sql.DB
}

func NewImageRepository(DB *sql.DB) ImageRepository {
	return ImageRepository{
		db: DB,
	}
}

func scanImage(s rowScanner) (models.Image, error) {
	image := models.Image{}

	err := s.Scan(&image.Id, &image.FileName, &image.FileSize, &image.MimeType)
	if err != nil {
		return image, err
	}

	image.URL = models.ImageURL(image.Id)

	return image, nil
}

func (repo ImageRepository) CreateImage(fileName string, fileSize int64, mimeType string) (int, error) {
	var id int

	err := repo.db.QueryRow(
		"INSERT INTO images (file_name, file_size, mime_type) VALUES ($1, $2, $3) RETURNING id",
		fileName, fileSize, mimeType,
	).Scan(&id)
	if err != nil {
		logger.Error.Printf("error insert image: %#v", err.Error())
		return 0, err
	}

	logger.Info.Printf("image created: %v", id)

	return id, nil
}

func (repo ImageRepository) GetImageByID(id int) (models.Image, error) {
	row := repo.db.QueryRow("SELECT id, file_name, file_size, mime_type FROM images WHERE id = $1", id)

	image, err := scanImage(row)
	if err != nil {
		logger.Error.Printf("error select image: %#v", err.Error())
		return image, err
	}

	return image, nil
}

func (repo ImageRepository) GetImagesByIDs(ids []int) ([]models.Image, error) {
	var images []models.Image

	rows, err := repo.db.Query("SELECT id, file_name, file_size, mime_type FROM images WHERE id = ANY($1)", pq.Array(ids))
	if err != nil {
		logger.Error.Printf("error select images: %#v", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		image, err := scanImage(rows)
		if err != nil {
			logger.Error.Printf("scan rows error: %#v", err.Error())
			return nil, err
		}

		images = append(images, image)
	}

	return images, nil
}

func (repo ImageRepository) ValidateImagesExist(ids []int) error {
	if len(ids) == 0 {
		return nil
	}

	images, err := repo.GetImagesByIDs(ids)
	if err != nil {
		return err
	}

	if len(images) != len(ids) {
		return models.ErrImageNotFound
	}

	return nil
}

func (repo ImageRepository) DeleteImage(id int) error {
	_, err := repo.db.Exec("DELETE FROM images WHERE id = $1", id)
	if err != nil {
		logger.Error.Printf("error delete image: %v", err.Error())
		return err
	}

	logger.Info.Printf("image deleted: %v", id)

	return nil
}
