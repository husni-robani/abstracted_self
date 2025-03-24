package repositories

import (
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
	ProjectURL string
	StartDate time.Time
	EndDate time.Time
	ImageId int
	Filename string
}

func (repo ProjectRepository) GetAllProjectsWithImages() ([]models.Project, error) {
	query := "SELECT p.id, p.name, p.description, p.tech_stack, p.source_url, p.project_url, p.start_date, p.end_date, pi.id, pi.file_name FROM projects p JOIN project_images pi ON p.id = pi.project_id"

	rows, err := repo.db.Query(query)
	if err != nil {
		logger.Error.Printf("failed exec query select: %v", err)
		return nil, err
	}

	projectContainers := []ProjectWithImage{}

	for rows.Next() {
		var project ProjectWithImage
		if err := rows.Scan(&project.ProjectId, &project.Name, &project.Description, pq.Array(&project.TechStack), pq.Array(&project.SourceURL), &project.ProjectURL, &project.StartDate, &project.EndDate, &project.ImageId, &project.Filename); err != nil {
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
			project.ProjectURL = projectContainer.ProjectURL
			project.StartDate = projectContainer.StartDate
			project.EndDate = projectContainer.EndDate
			project.Images = append(project.Images, models.ProjectImage{
				Id: projectContainer.ImageId,
				FileName: projectContainer.Filename,
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
			project.ProjectURL = projectContainer.ProjectURL
			project.StartDate = projectContainer.StartDate
			project.EndDate = projectContainer.EndDate
			project.Images = append(project.Images, models.ProjectImage{
				Id: projectContainer.ImageId,
				FileName: projectContainer.Filename,
			})

			projects = append(projects, project)
			continue
		}

		if projects[len(projects) - 1].Id == projectContainer.ProjectId {
			projects[len(projects) - 1].Images = append(projects[len(projects)-1].Images, models.ProjectImage{
				Id: projectContainer.ImageId,
				FileName: projectContainer.Filename,
			})
		}
	}

	// for i, project := range projects {
	// 	logger.Info.Printf("%d: %#v", i, project)
	// }

	return projects, nil
}