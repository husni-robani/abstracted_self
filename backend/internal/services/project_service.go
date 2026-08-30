package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
)

type ProjectService struct {
	projectRepo repositories.ProjectRepository
	imageRepo   repositories.ImageRepository
}

func NewProjectService(projectRepo repositories.ProjectRepository, imageRepo repositories.ImageRepository) ProjectService {
	return ProjectService{
		projectRepo: projectRepo,
		imageRepo:   imageRepo,
	}
}

func (service ProjectService) CreateNewProject(project_data requests.CreateProjectRequest) error {
	// insert project data to database
	projectId, err := service.projectRepo.CreateNewProject(project_data);
	if  err != nil {
		return err
	}

	// store each image to storage and images table
	imageIds := make([]int, 0, len(project_data.Images))
	for i := range project_data.Images {
		extension := filepath.Ext(project_data.Images[i].Filename)
		newFileName := uuid.New().String() + extension
		mimeType := project_data.Images[i].Header.Get("Content-Type")

		project_data.Images[i].Filename = newFileName

		if err := utils.SaveFile(&project_data.Images[i], "." + os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
			return err
		}

		imageId, err := service.imageRepo.CreateImage(newFileName, project_data.Images[i].Size, mimeType)
		if err != nil {
			return err
		}
		imageIds = append(imageIds, imageId)
	}

	// insert project_images mappings
	if err := service.projectRepo.AddProjectImages(projectId, imageIds); err != nil {
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

	images, err := service.projectRepo.GetImagesByProjectId(id)
	if err != nil {
		return models.Project{}, err
	}

	project.Images = images

	return project, nil
}

func (service ProjectService) DeleteProjectById(id int) error {
	images, err := service.projectRepo.GetImagesByProjectId(id)
	if err != nil {
		return err
	}

	if err := service.projectRepo.DeleteProjectById(id); err != nil {
		return err
	}

	for _, image := range images {
		if err := service.imageRepo.DeleteImage(image.Id); err != nil {
			logger.Error.Printf("delete image row failed: %v", err.Error())
			continue
		}

		if err := utils.RemoveFile("." + os.Getenv("IMAGES_STORAGE_PATH") + "/", image.FileName); err != nil {
			logger.Error.Printf("delete image file from storage failed: %v", err.Error())
		}
	}

	return nil
}

func (service ProjectService) UpdateProject(id int, project requests.UpdateProjectRequest) error {
	// Handle NewImages and ImagesToDelete differently

	// []TechStack data processing for query generation
	techStack := "{}"
	if len(project.TechStack) > 0 {
		techStack = "{"
		for i, v := range project.TechStack {
			if i == len(project.TechStack) - 1{
				techStack += fmt.Sprintf("%v }", v)
				break
			}

			techStack += fmt.Sprintf("%v, ", v)
		}
	}

	// []SourceURL data processing for query generation
	sourceUrl := "{}"
	if len(project.SourceURL) > 0 {
		sourceUrl = "{"
		for i, v := range project.SourceURL {
			if i == len(project.SourceURL) - 1{
				sourceUrl += fmt.Sprintf("%v }", v)
				break
			}

			sourceUrl += fmt.Sprintf("%v, ", v)
		}
	}


	fieldsUpdate := map[string]any{
		"name": project.Name,
		"description": project.Description,
		"tech_stack": techStack,
		"source_url": sourceUrl,
		"project_url": project.ProjectURL,
		"start_date": project.StartDate,
		"end_date": project.EndDate,
	}

	query := utils.GenerateSingleUpdateQuery("projects", fieldsUpdate, fmt.Sprintf("WHERE id = %d;", id))

	if err := service.projectRepo.UpdateProjectById(query); err != nil {
		return err
	}

	// Handle NewImages
	if len(project.NewImages) > 0 {
		imageIds := make([]int, 0, len(project.NewImages))
		for i := range project.NewImages {
			extension := filepath.Ext(project.NewImages[i].Filename)
			newFileName := uuid.New().String() + extension
			mimeType := project.NewImages[i].Header.Get("Content-Type")

			project.NewImages[i].Filename = newFileName

			if err := utils.SaveFile(&project.NewImages[i], "." + os.Getenv("IMAGES_STORAGE_PATH")); err != nil {
				return err
			}

			imageId, err := service.imageRepo.CreateImage(newFileName, project.NewImages[i].Size, mimeType)
			if err != nil {
				return err
			}
			imageIds = append(imageIds, imageId)
		}

		if err := service.projectRepo.AddProjectImages(id, imageIds); err != nil {
			return err
		}
	}

	// Handle ImagesToDelete (image ids)
	for _, imageId := range project.ImagesToDelete {
		image, err := service.imageRepo.GetImageByID(imageId)
		if err != nil {
			logger.Error.Printf("get image to delete failed: %v", err.Error())
			continue
		}

		if err := service.imageRepo.DeleteImage(image.Id); err != nil {
			logger.Error.Printf("delete image row failed: %v", err.Error())
			continue
		}

		if err := utils.RemoveFile("." + os.Getenv("IMAGES_STORAGE_PATH") + "/", image.FileName); err != nil {
			logger.Error.Printf("delete image file from storage failed: %v", err.Error())
		}
	}

	return nil
}
