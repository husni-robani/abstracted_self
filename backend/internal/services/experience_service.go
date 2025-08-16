package services

import (
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type ExperienceService struct {
	repo repositories.ExperienceRepo
}

func NewExperienceService(repo repositories.ExperienceRepo) ExperienceService {
	return ExperienceService{
		repo: repo,
	}
}

func (service ExperienceService) GetAllExperiences() ([]models.Experience, error) {
	experiences, err := service.repo.GetAllExperiences()
	if err != nil {
		return nil, err
	}

	return experiences, nil 
}

func (service ExperienceService) CreateExperience(experience requests.CreateExperienceRequest) error {
	if err := service.repo.CreateExperience(experience); err != nil {
		return err
	}

	return nil
}

func (service ExperienceService) UpdateExperience(id int, experience requests.UpdateExperienceRequest) error {
	var values = make(map[string]any)

	if experience.JobTitle != nil {
		values["job_title"] = *experience.JobTitle
	}

	if experience.CompanyName != nil {
		values["company_name"] = *experience.CompanyName
	}

	if experience.WorkPlace != nil {
		values["work_place"] = *experience.WorkPlace
	}

	if experience.StartDate != nil {
		values["start_date"] = *experience.StartDate
	}

	if experience.EndDate != nil {
		if *experience.EndDate == "" {
			values["end_date"] = nil
		}else {
			values["end_date"] = *experience.EndDate
		}
	}

	if experience.Description != nil {
		values["description"] = *experience.Description
	}

	// accomplishments
	if experience.Accomplishments != nil {
		accomplishments := "{"
		for i, v := range experience.Accomplishments {
			if i == len(experience.Accomplishments) - 1{
				accomplishments += fmt.Sprintf("%v }", v)
				break
			}

			accomplishments += fmt.Sprintf("%v, ", v)
		}

		values["accomplishments"] = accomplishments
	}

	// tech_stack
	if experience.TechStack != nil {
		tech_stack := "{"
		for i, v := range experience.TechStack {
			if i == len(experience.TechStack) - 1{
				tech_stack += fmt.Sprintf("%v }", v)
				break
			}
	
			tech_stack += fmt.Sprintf("%v, ", v)
		}
	
		values["tech_stack"] = tech_stack
	}


	// update to database
	if err := service.repo.UpdateExperience(id, values); err != nil {
		return err
	}

	return nil
}

func (service ExperienceService) GetExperienceById(id int) (models.Experience, error){
	experience, err := service.repo.GetExperienceById(id)
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (service ExperienceService) DeleteExperience(id int) error {
	if err := service.repo.DeleteExperience(id); err != nil {
		return err
	}

	return nil
}