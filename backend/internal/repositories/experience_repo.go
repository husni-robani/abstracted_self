package repositories

import (
	"database/sql"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/lib/pq"
)

type ExperienceRepo struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) ExperienceRepo {
	return ExperienceRepo{
		db: db,
	}
}

func (repo ExperienceRepo) GetAllExperiences() ([]models.Experience, error) {
	query := "SELECT id, job_title, company_name, work_place, start_date, end_date, description, accomplishments from experiences"

	rows, err := repo.db.Query(query)
	if err != nil {
		logger.Error.Printf("failed query select to experiences: %v", err)
		return nil, err
	}

	experiences := []models.Experience{}
	for rows.Next() {
		experience := models.Experience{}

		err := rows.Scan(&experience.Id, &experience.JobTitle, &experience.CompanyName, &experience.WorkPlace, &experience.StartDate, &experience.EndDate, &experience.Description, pq.Array(&experience.Accomplishments))

		if err != nil {
			logger.Error.Printf("failed scan row: %v", err)
			return nil, err
		}

		experiences = append(experiences, experience)
	}

	return experiences, nil
}

func (repo ExperienceRepo) CreateExperience(experience requests.CreateExperienceRequest) error {
	result, err := repo.db.Exec("INSERT INTO experiences (job_title, company_name, work_place, start_date, end_date, description, accomplishments) VALUES ($1, $2, $3, $4, $5, $6, $7)", experience.JobTitle, experience.CompanyName, experience.WorkPlace, experience.StartDate, experience.EndDate, experience.Description, pq.Array(experience.Accomplishments))

	if err != nil {
		logger.Error.Printf("insert experience failed: %v", err)
		return err
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed get total row affected: %v", err)
		return err
	}

	logger.Info.Printf("Rows affected: %v", rowAffected)


	return nil
}