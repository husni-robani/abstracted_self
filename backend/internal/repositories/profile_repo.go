package repositories

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type ProfileRepository struct {}

func NewProfileRepository() (ProfileRepository) {
	return ProfileRepository{}
}

func (ProfileRepository) ReadProfileData() (models.Profile, error) {
	bytes, err := os.ReadFile(os.Getenv("PROFILE_DB_PATH"))
	if err != nil {
		logger.Error.Printf("Failed to read json file: %v", err)
		return models.Profile{}, err
	}

	var profileData models.Profile
	if err := json.Unmarshal(bytes, &profileData); err != nil {
		logger.Error.Printf("Failed to unmarshal profile data: %v", err)
		return models.Profile{}, err
	}

	return profileData, nil
}

func (ProfileRepository) WriteProfileData(newProfile models.Profile) (error) {
	marshaledData, err := json.MarshalIndent(newProfile, "", " ")
	if err != nil {
		logger.Error.Println("failed to marshal: ", err)
		return errors.New("failed to marshal profile data")
	}

	if err := os.WriteFile(os.Getenv("PROFILE_DB_PATH"), marshaledData, 0644); err != nil {
		logger.Error.Println("failed to write file: ", err)
		return errors.New("failed to write profile data")
	}
	return nil
}

func (repo ProfileRepository) GetAllSkillTypeName() ([]string, error) {
	profile, err := repo.ReadProfileData()
	if err != nil {
		return nil, err
	}

	var result []string

	for _, skillType := range profile.SkillSet{
		result = append(result, skillType.TypeName)
	}

	return result, nil
}