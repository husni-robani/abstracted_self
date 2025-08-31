package services

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
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
		err := utils.SaveFile(&project_data.Images[i], "." + os.Getenv("IMAGES_STORAGE_PATH"))
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

func (service ProjectService) DeleteProjectById(id int) error {
	if err := service.projectRepo.DeleteProjectById(id); err != nil {
		return err
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
		// set new filename and save to storage
		for i := range project.NewImages {
			// generate new filename
			extension := filepath.Ext(project.NewImages[i].Filename)
			newFileName := uuid.New().String() + extension
			
			// set new filename to image
			project.NewImages[i].Filename = newFileName

			// store to storage
			err := utils.SaveFile(&project.NewImages[i], "." + os.Getenv("IMAGES_STORAGE_PATH"))
			if err != nil {
				return err
			}
		}
		
		err := service.projectImageRepo.AddProjectImages(id, project.NewImages);
		if err != nil {
			return err
		}
	}

	// Handle ImagesToDelete
	if len(project.ImagesToDelete) > 0 {
		// delete the file in storage
		var wg sync.WaitGroup
		for _, id := range project.ImagesToDelete {
			wg.Add(1)
			go func(){
				defer wg.Done()
				projectImage, err := service.projectImageRepo.GetImageById(id)
				if err != nil {
					logger.Error.Printf("delete image file from storage failed: %v", err)
					return
				}

				if err := utils.RemoveFile("." + os.Getenv("IMAGES_STORAGE_PATH") + "/", projectImage.FileName); err != nil {
					logger.Error.Printf("delete image file from storage failed: %v", err)
					return
				}

				logger.Info.Printf("Delete %s from storage succeeded", projectImage.FileName)
			}()
		}
		wg.Wait()

		// delete project_images data
		err := service.projectImageRepo.DeleteProjectImageByID(project.ImagesToDelete...)
		if err != nil {
			return err
		}
	}

	return nil
}