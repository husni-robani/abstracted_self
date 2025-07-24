package repositories

import (
	"encoding/json"
	"os"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type ProfileRepository struct {}

func NewProfileRepository() (ProfileRepository) {
	return ProfileRepository{}
}

func (repository ProfileRepository) GetName() (string, error) {
	bytes, err := os.ReadFile("./internal/db/profile_data.json")
	if err != nil {
		logger.Error.Printf("Failed to read json file: %v", err)
		return "", err
	}

	var profileData models.Profile
	if err := json.Unmarshal(bytes, &profileData); err != nil {
		logger.Error.Printf("Failed to unmarshal profile data: %v", err)
		return "", err
	}

	return profileData.Name, nil
}