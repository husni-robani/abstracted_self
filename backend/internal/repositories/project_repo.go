package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
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

type ProjectWithImage struct {
	ProjectId int
	Name string
	Description string
	TechStack []string
	SourceURL []string
	ProjectURL sql.NullString
	StartDate sql.NullTime
	EndDate sql.NullTime
	ImageId sql.NullInt32
	FileName sql.NullString
}

func (repo ProjectRepository) GetAllProjectsWithImages() ([]models.Project, error) {
	query := "SELECT p.id, p.name, p.description, p.tech_stack, p.source_url, p.project_url, p.start_date, p.end_date, pi.id, pi.file_name FROM projects p LEFT JOIN project_images pi ON p.id = pi.project_id ORDER BY p.id desc"

	rows, err := repo.db.Query(query)
	if err != nil {
		logger.Error.Printf("failed exec query select: %v", err)
		return nil, err
	}

	projectContainers := []ProjectWithImage{}

	for rows.Next() {
		var project ProjectWithImage
		if err := rows.Scan(&project.ProjectId, &project.Name, &project.Description, pq.Array(&project.TechStack), pq.Array(&project.SourceURL), &project.ProjectURL, &project.StartDate, &project.EndDate, &project.ImageId, &project.FileName); err != nil {
			logger.Error.Printf("failed to scan project image: %v", err)
			return nil, err
		}

		projectContainers = append(projectContainers, project)
	}

	var projects []models.Project

	for _, projectContainer := range projectContainers {

		project := models.Project{}
		if len(projects) == 0{
			project.Id = projectContainer.ProjectId
			project.Name = projectContainer.Name
			project.Description = projectContainer.Description
			project.TechStack = projectContainer.TechStack
			project.SourceURL = projectContainer.SourceURL
			project.ProjectURL = projectContainer.ProjectURL.String
			project.StartDate = projectContainer.StartDate.Time
			project.EndDate = projectContainer.EndDate.Time
			project.Images = append(project.Images, models.ProjectImage{
				Id: int(projectContainer.ImageId.Int32),
				FileName: projectContainer.FileName.String,
			})

			projects = append(projects, project)
			continue
		}

		if projects[len(projects) - 1].Id != projectContainer.ProjectId {
			project.Id = projectContainer.ProjectId
			project.Name = projectContainer.Name
			project.Description = projectContainer.Description
			project.TechStack = projectContainer.TechStack
			project.SourceURL = projectContainer.SourceURL
			project.ProjectURL = projectContainer.ProjectURL.String
			project.StartDate = projectContainer.StartDate.Time
			project.EndDate = projectContainer.EndDate.Time
			project.Images = append(project.Images, models.ProjectImage{
				Id: int(projectContainer.ImageId.Int32),
				FileName: projectContainer.FileName.String,
			})

			projects = append(projects, project)
			continue
		}

		if projects[len(projects) - 1].Id == projectContainer.ProjectId {
			projects[len(projects) - 1].Images = append(projects[len(projects)-1].Images, models.ProjectImage{
				Id: int(projectContainer.ImageId.Int32),
				FileName: projectContainer.FileName.String,
			})
		}
	}

	return projects, nil
}

func (repo ProjectRepository) GetProjectById(id int) (models.Project, error) {
	var project models.Project

	row := repo.db.QueryRow("SELECT p.id, p.name, p.description, p.tech_stack, p.source_url, p.project_url, p.start_date, p.end_date FROM projects p WHERE id = $1", id)

	if err := row.Scan(&project.Id, &project.Name, &project.Description, pq.Array(&project.TechStack), pq.Array(&project.SourceURL), &project.ProjectURL, &project.StartDate, &project.EndDate); err != nil {
		logger.Error.Printf("failed to scan: %v", err)
		return models.Project{}, nil
	}

	return project, nil
}

func (repo ProjectRepository) DeleteProjectById(id int) error {
	result, err := repo.db.Exec("DELETE FROM projects WHERE id = $1", id)
	if err != nil {
		logger.Error.Printf("failed to exec query delete: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed to get rows affected: %v", err)
		return err
	}

	logger.Info.Printf("rows affected: %v", rowsAffected)

	return nil
}

func (repo ProjectRepository) UpdateProjectById(query string) error {
	timeOut, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()
	
	result, err := repo.db.ExecContext(timeOut, query)
	if err != nil {
		logger.Error.Printf("Query execution is failed: %v", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	logger.Info.Printf("update project is succeeded | Rows affected: %d", rowsAffected)

	return nil
}