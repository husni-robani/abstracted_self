package services

import (
	"path/filepath"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type ProjectService struct {
	projectRepo repositories.ProjectRepository
	projectImageRepo repositories.ProjectImageRepository
}

func NewProjectService(projectRepo repositories.ProjectRepository, projectImageRepo repositories.ProjectImageRepository) ProjectService {
	return ProjectService{
		projectRepo: projectRepo,
		projectImageRepo: projectImageRepo,
	}
}

func (service ProjectService) CreateNewProject(project_data requests.CreateProjectRequest) error {
	// insert project data to database
	projectId, err := service.projectRepo.CreateNewProject(project_data);
	if  err != nil {
		return err
	}

	// set new filename and save to storage
	for i := range project_data.Images {
		// generate new filename
		extension := filepath.Ext(project_data.Images[i].Filename)
		newFileName := uuid.New().String() + extension
		
		// set new filename to image
		project_data.Images[i].Filename = newFileName

		// store to storage
		err := utils.SaveFile(&project_data.Images[i], "./storage/project_images")
		if err != nil {
			return err
		}
	}

	// insert images data to database
	if err := service.projectImageRepo.AddProjectImages(projectId, project_data.Images); err != nil {
		return err
	}
	
	return nil
}

func (service ProjectService) GetAllProjectsWithImages() ([]models.Project, error){
	projects, err := service.projectRepo.GetAllProjectsWithImages()
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (service ProjectService) GetProjectById(id int) (models.Project, error) {
	project, err := service.projectRepo.GetProjectById(id)
	if err != nil {
		return models.Project{}, err
	}

	images, err := service.projectImageRepo.GetImagesByProjectId(id)
	if err != nil {
		return models.Project{}, err
	}

	project.Images = images

	return project, nil
}