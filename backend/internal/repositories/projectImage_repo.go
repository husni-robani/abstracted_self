package repositories

import "database/sql"

type ProjectImageRepository struct {
	db *sql.DB
}

func NewProjectImageRepository(db *sql.DB) ProjectImageRepository {
	return ProjectImageRepository{db: db}
}