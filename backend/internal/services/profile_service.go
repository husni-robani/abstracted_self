package services

import (
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
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
	ResumeFileName string `json:"resume_file_name,omitempty"`
	SkillSet []models.SkillType `json:"skill_set,omitempty"`
}

func (service ProfileService) GetProfileData(name bool, summary bool, bio bool, taglines bool, resume bool, skill_set bool) (ProfileData, error) {
	var dataResult ProfileData
	
	profileData, err := service.Repository.ReadProfileData()
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
		dataResult.ResumeFileName = profileData.ResumeFileName
	}
	if skill_set {
		dataResult.SkillSet = profileData.SkillSet
	}


	return dataResult, nil
}

func (service ProfileService) UpdateProfileData(dataRequest requests.UpdateProfileRequest) (error) {
	originalProfile, err := service.Repository.ReadProfileData()
	if err != nil {
		return err
	}

	// check request consist of new resume file
	if dataRequest.ResumeFile != nil {
		// save the file
		newFileName, err := saveNewResume(dataRequest.ResumeFile)
		if err != nil {
			return err
		}
		// delete old resume file
		utils.RemoveFile("./storage/documents/", originalProfile.ResumeFileName)
		// set the new resume_file_name
		originalProfile.ResumeFileName = newFileName
	}

	// check which attributes are not nill value
	if dataRequest.Name != "" {
		originalProfile.Name = dataRequest.Name
	}
	if dataRequest.Bio != "" {
		originalProfile.Bio = dataRequest.Bio
	}
	if dataRequest.Summary != "" {
		originalProfile.Summary = dataRequest.Summary
	}
	if dataRequest.Taglines != nil {
		originalProfile.Taglines = dataRequest.Taglines
	}

	// save new profile
	if err := service.Repository.WriteProfileData(originalProfile); err != nil {
		return err
	}

	return nil
}


// helper function for UpdateProfileData, returns new filename and error
func saveNewResume(resumeFile *multipart.FileHeader) (file_name string, err error){
	// Generate file name
	extension := filepath.Ext(resumeFile.Filename)
	newFileName := uuid.New().String() + extension

	// set new file name
	resumeFile.Filename = newFileName

	// store to storage
	if err := utils.SaveFile(resumeFile, "./storage/documents"); err != nil {
		logger.Error.Printf("store resume file to storage failed: %v", err)
		return "", err
	}

	return resumeFile.Filename, nil
}

func (service ProfileService) AddSkillSetType(dataRequest requests.AddProfileSkillSetType) (error) {
	if err := service.Repository.AddSkillSetType(dataRequest.TypeName); err != nil {
		return err
	}
	return nil
}

func (service ProfileService) AddSkill(dataRequest requests.AddProfileSkill) (error) {
	// insert icon
		// Generate file name
	extension := filepath.Ext(dataRequest.IconFile.Filename)
	newFileName := uuid.New().String() + extension
		// set new file name
	dataRequest.IconFile.Filename = newFileName

	// store to storage
	if err := utils.SaveFile(dataRequest.IconFile, "./storage/icons"); err != nil {
		logger.Error.Printf("store resume file to storage failed: %v", err)
		return err
	}
	
	// add skill
	newSkill := models.Skill{
		Name: dataRequest.Name,
		IconFilename: dataRequest.IconFile.Filename,
		IsMostUsed: *dataRequest.IsMostUsed,
		// IsMostUsed: false,
	}
	if err := service.Repository.AddSkill(newSkill, dataRequest.SkillType); err != nil {
		return err
	}
	return nil
}

func (service ProfileService) RemoveSkillType(typeName string) (error) {
	if err := service.Repository.RemoveSkillType(typeName); err != nil {
		return err
	}

	return nil
}

func (service ProfileService) RemoveSkill(skillName string) (error) {
	if err := service.Repository.RemoveSkill(skillName); err != nil {
		return err
	}

	return nil
}

func (service ProfileService) ToggleIsMostUsed(skillName string, typeName string) error {
	if err := service.Repository.ToggleIsMostUsed(skillName, typeName); err != nil {
		return err
	}

	return nil
}