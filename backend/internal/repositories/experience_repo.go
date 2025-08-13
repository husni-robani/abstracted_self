package repositories

import (
	"database/sql"
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
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
	query := "SELECT id, job_title, company_name, work_place, start_date, end_date, description, accomplishments, tech_stack FROM experiences ORDER BY start_date desc"

	rows, err := repo.db.Query(query)
	if err != nil {
		logger.Error.Printf("failed query select to experiences: %v", err)
		return nil, err
	}

	experiences := []models.Experience{}
	for rows.Next() {
		experience := models.Experience{}

		err := rows.Scan(&experience.Id, &experience.JobTitle, &experience.CompanyName, &experience.WorkPlace, &experience.StartDate, &experience.EndDate, &experience.Description, pq.Array(&experience.Accomplishments), pq.Array(&experience.TechStack))

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

func (repo ExperienceRepo) UpdateExperience(id int, experience map[string]any) error{
	query := utils.GenerateSingleUpdateQuery("experiences", experience, fmt.Sprintf("where id = %v", id))

	result, err := repo.db.Exec(query)
	if err != nil {
		logger.Error.Printf("failed exec query update: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed to get total rows affected")
		return err
	}

	logger.Info.Printf("rows affected: %v", rowsAffected)

	return nil 
}

func (repo ExperienceRepo) GetExperienceById(id int) ( models.Experience, error) {
	row := repo.db.QueryRow("SELECT id, job_title, company_name, work_place, start_date, end_date, description, accomplishments FROM experiences WHERE id = $1", id)

	var experience models.Experience
	err := row.Scan(&experience.Id, &experience.JobTitle, &experience.CompanyName, &experience.WorkPlace, &experience.StartDate, &experience.EndDate, &experience.Description, pq.Array(&experience.Accomplishments))

	if err != nil {
		logger.Error.Printf("failed scan row to experience model: %v", err)
		return experience, err
	}

	return experience, nil
}

func (repo ExperienceRepo) DeleteExperience(id int) error {
	result, err := repo.db.Exec("DELETE FROM experiences WHERE id = $1", id)
	if err != nil {
		logger.Error.Printf("failed exec query delete experience: %v", err.Error())
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed to get total of rows affected: %v", err.Error())
		return err
	}

	logger.Info.Printf("rows affected: %v", rowsAffected)

	return nil
}