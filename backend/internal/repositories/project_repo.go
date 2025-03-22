package repositories

import (
	"database/sql"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/lib/pq"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return ProjectRepository{db: db}
}

func (repo ProjectRepository) CreateNewProject(project requests.CreateProjectRequest) (projectId int, err error) {
	err = repo.db.QueryRow("INSERT INTO projects (name, description, tech_stack, source_url, project_url, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", project.Name, project.Description, pq.Array(project.TechStack), pq.Array(project.SourceURL), project.ProjectURL, project.StartDate, project.EndDate).Scan(&projectId)

	if err != nil {
		logger.Error.Printf("failed to insert new project: %v", err)
		return 0, err
	}

	logger.Info.Println("create new project successful!")

	return int(projectId), nil
}