package repositories

import (
	"database/sql"
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type TechnologyRepository struct {
	db *sql.DB
}

func NewTechnologyRepository(db *sql.DB) TechnologyRepository {
	return TechnologyRepository{
		db: db,
	}
}

func (repo TechnologyRepository) GetAllTechnologies() ([]models.Technology, error) {
	var technologies []models.Technology

	rows, err := repo.db.Query("SELECT id, type_id, name FROM technologies")
	if err != nil {
		logger.Error.Printf("failed select technologies: %v", err)
		return nil, err
	}

	for rows.Next() {
		var tech models.Technology
		if err := rows.Scan(&tech.Id, &tech.TypeId, &tech.Name); err != nil {
			logger.Error.Printf("failed scan: %v", err)
			return nil, err
		}

		technologies = append(technologies, tech)
	}

	return technologies, nil
}

func (repo TechnologyRepository) IsTypeExist(id int) (bool, error) {
	var exist bool

	row := repo.db.QueryRow("SELECT EXISTS(SELECT FROM technology_types WHERE id = $1)", id)

	if err := row.Scan(&exist); err != nil {
		logger.Error.Printf("failed scan row: %v", err)
		return false, err
	}

	return exist, nil
}

func (repo TechnologyRepository) CreateTechnologies(technologies []requests.CreateTechnologyRequest) error {
	query := "INSERT INTO technologies (type_id, name) values "

	for i, tech := range technologies {
		if i == len(technologies) - 1 {
			query += fmt.Sprintf("('%d', '%s');", tech.TypeId, tech.Name)
		}else {
			query += fmt.Sprintf("('%d', '%s'), ", tech.TypeId, tech.Name)
		}
	}

	result, err := repo.db.Exec(query)
	if err != nil {
		logger.Error.Printf("failed insert technologies: %v", err)
		return err
	}

	totalRowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed get rows affected: %v", err)
		return err
	}

	logger.Info.Printf("rows affected: %v", totalRowsAffected)

	return nil
}

func (repo TechnologyRepository) DeleteTechnologyByID(id int) error {
	query := "DELETE FROM technologies WHERE id = $1"

	result, err := repo.db.Exec(query, id)
	if err != nil {
		logger.Error.Printf("failed exec query delete technology: %v", err)
		return err
	}

	totalRowAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed get rows effected: %v", err)
		return err
	}

	logger.Info.Printf("rows affected: %v", totalRowAffected)

	return nil
}