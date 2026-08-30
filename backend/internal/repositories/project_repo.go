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

func scanProject(row rowScanner, project *models.Project) error {
	var projectURL sql.NullString
	var startDate sql.NullTime
	var endDate sql.NullTime

	if err := row.Scan(&project.Id, &project.Name, &project.Description, pq.Array(&project.TechStack), pq.Array(&project.SourceURL), &projectURL, &startDate, &endDate); err != nil {
		return err
	}

	project.ProjectURL = projectURL.String
	if startDate.Valid {
		project.StartDate = startDate.Time
	}
	if endDate.Valid {
		project.EndDate = endDate.Time
	}

	return nil
}

const projectSelect = "SELECT id, name, description, tech_stack, source_url, project_url, start_date, end_date FROM projects"

func (repo ProjectRepository) GetAllProjectsWithImages() ([]models.Project, error) {
	rows, err := repo.db.Query(projectSelect + " ORDER BY id desc")
	if err != nil {
		logger.Error.Printf("failed exec query select: %v", err)
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		project := models.Project{}
		if err := scanProject(rows, &project); err != nil {
			logger.Error.Printf("failed to scan project: %v", err)
			return nil, err
		}

		projects = append(projects, project)
	}

	imageRows, err := repo.db.Query(
		"SELECT pi.project_id, i.id, i.file_name, i.file_size, i.mime_type FROM project_images pi JOIN images i ON i.id = pi.image_id ORDER BY i.id",
	)
	if err != nil {
		logger.Error.Printf("failed exec query select project images: %v", err)
		return nil, err
	}
	defer imageRows.Close()

	imagesByProject := make(map[int][]models.Image)

	for imageRows.Next() {
		var projectId int
		image := models.Image{}

		if err := imageRows.Scan(&projectId, &image.Id, &image.FileName, &image.FileSize, &image.MimeType); err != nil {
			logger.Error.Printf("failed to scan project image: %v", err)
			return nil, err
		}

		image.URL = models.ImageURL(image.Id)
		imagesByProject[projectId] = append(imagesByProject[projectId], image)
	}

	for i := range projects {
		projects[i].Images = imagesByProject[projects[i].Id]
	}

	return projects, nil
}

func (repo ProjectRepository) GetProjectById(id int) (models.Project, error) {
	var project models.Project

	row := repo.db.QueryRow(projectSelect+" WHERE id = $1", id)

	if err := scanProject(row, &project); err != nil {
		logger.Error.Printf("failed to scan: %v", err)
		return models.Project{}, err
	}

	return project, nil
}

func (repo ProjectRepository) GetImagesByProjectId(projectId int) ([]models.Image, error) {
	rows, err := repo.db.Query(
		"SELECT i.id, i.file_name, i.file_size, i.mime_type FROM project_images pi JOIN images i ON i.id = pi.image_id WHERE pi.project_id = $1 ORDER BY i.id",
		projectId,
	)
	if err != nil {
		logger.Error.Printf("failed to exec query select project images: %v", err)
		return nil, err
	}
	defer rows.Close()

	var images []models.Image

	for rows.Next() {
		image := models.Image{}

		if err := rows.Scan(&image.Id, &image.FileName, &image.FileSize, &image.MimeType); err != nil {
			logger.Error.Printf("failed to scan data: %v", err)
			return nil, err
		}

		image.URL = models.ImageURL(image.Id)
		images = append(images, image)
	}

	return images, nil
}

func (repo ProjectRepository) AddProjectImages(projectId int, imageIds []int) error {
	for _, imageId := range imageIds {
		if _, err := repo.db.Exec("INSERT INTO project_images (project_id, image_id) VALUES ($1, $2)", projectId, imageId); err != nil {
			logger.Error.Printf("failed to insert project_image mapping to database: %v", err)
			return err
		}
	}

	logger.Info.Printf("create project images mappings successful | project_id: %d", projectId)

	return nil
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
