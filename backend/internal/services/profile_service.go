package services

import (
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type ProfileService struct {
	Repository repositories.ProfileRepository
}

func NewProfileService(repo repositories.ProfileRepository) (ProfileService) {
	return ProfileService{
		Repository: repo,
	}
}

type ProfileData struct {
	Name string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Bio string `json:"bio,omitempty"`
	Taglines []string `json:"taglines,omitempty"`
	ResumeFilePath string `json:"resume_file_path,omitempty"`
}

func (service ProfileService) GetProfileData(name bool, summary bool, bio bool, taglines bool, resume bool) (ProfileData, error) {
	var dataResult ProfileData
	
	profileData, err := service.Repository.GetProfileData()
	if err != nil {
		logger.Error.Println("Failed to get profile data")
		return dataResult, err
	}

	if name {
		dataResult.Name = profileData.Name
	}
	if summary {
		dataResult.Summary = profileData.Summary
	}
	if bio {
		dataResult.Bio = profileData.Bio
	}
	if taglines {
		dataResult.Taglines = profileData.Taglines
	}
	if resume {
		dataResult.ResumeFilePath = profileData.ResumeFilePath
	}


	return dataResult, nil
}