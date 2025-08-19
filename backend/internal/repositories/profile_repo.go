package repositories

import (
	"encoding/json"
	"errors"
	"os"
	"slices"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/utils"
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

func (repo ProfileRepository) GetAllSkillName() ([]string, error){
	profile, err := repo.ReadProfileData()
	if err != nil {
		return nil, err
	}

	var skillNames []string

	for _, skillType := range profile.SkillSet {
		for _, skill := range skillType.SkillItems {
			skillNames = append(skillNames, skill.Name)
		}
	}

	return skillNames, nil
}

func (repo ProfileRepository) AddSkillSetType(typeName string) (error) {
	// Check is the type_name already exists
	typeNames, err := repo.GetAllSkillTypeName()
	if err != nil {
		return err
	}
	if slices.Contains(typeNames, typeName) {
		return models.ErrSkillTypeDuplicate
	}

	// append the new type name to SkillSet
	originalProfileData, err := repo.ReadProfileData()
	if err != nil {
		return err
	}

	originalProfileData.SkillSet = append(originalProfileData.SkillSet, models.SkillType{TypeName: typeName, SkillItems: []models.Skill{}})

	if err := repo.WriteProfileData(originalProfileData); err != nil {
		return err
	}

	return nil
}

func (repo ProfileRepository) AddSkill(newSkill models.Skill, typeName string) (error){
	// check is type already exists
	typeNames, err := repo.GetAllSkillTypeName()
	if err != nil {
		return err
	}
	if !slices.Contains(typeNames, typeName) {
		logger.Info.Printf("skill type not found!")
		return models.ErrSkillTypeNotFound
	}

	// check is skill already exists
	skillNames, err := repo.GetAllSkillName()
	if err != nil {
		return err
	}
	if slices.Contains(skillNames, newSkill.Name){
		logger.Info.Printf("skill already exists!")
		return models.ErrSkillDuplicate
	}

	// Get profile data and append new skill
	profile, err := repo.ReadProfileData()
	if err != nil {
		return err
	}

	for i, skillType := range profile.SkillSet {
		if skillType.TypeName == typeName {
			profile.SkillSet[i].SkillItems = append(skillType.SkillItems, newSkill)
		}
	}

	if err := repo.WriteProfileData(profile); err != nil {
		return err
	}

	return nil
}

func (repo ProfileRepository) RemoveSkillType(typeName string) (error) {
	allTypeNames, err := repo.GetAllSkillTypeName()
	if err != nil {
		return err
	}
	if !slices.Contains(allTypeNames, typeName) {
		return models.ErrSkillTypeNotFound
	}

	// get profile data 
	profileData, err := repo.ReadProfileData()
	if err != nil {
		return err
	}

	// remove skill type and related skill icons
	for i, skillType := range profileData.SkillSet {
		if skillType.TypeName == typeName {
			// remove all related skill icons
			for _, skill := range skillType.SkillItems {
				if err := utils.RemoveFile("./storage/icons/", skill.IconFilename); err != nil {
					logger.Error.Printf("remove file (%s) failed: %v", skill.IconFilename, err)
				}
			}
			// remove skill type
			profileData.SkillSet = slices.Delete(profileData.SkillSet, i, i+1)
		}
	}

	if err := repo.WriteProfileData(profileData); err != nil {
		return err
	}

	return nil
}

func (repo ProfileRepository) RemoveSkill(skillName string) error {
	profileData, err := repo.ReadProfileData()
	if err != nil {
		return err
	}

	isExists := false
	for i, skillType := range profileData.SkillSet{
		for k, skill := range skillType.SkillItems{
			if skill.Name == skillName {
				// remove the skill icon
				if err := utils.RemoveFile("./storage/icons/", skill.IconFilename); err != nil {
					logger.Error.Printf("remove file (%s) failed: %v", skill.IconFilename, err)
				}
				// remove the skill
				profileData.SkillSet[i].SkillItems = slices.Delete(profileData.SkillSet[i].SkillItems, k, k+1)
				isExists = true
			}
		}
	}

	if !isExists {
		return models.ErrSkillNotFound
	}

	if err := repo.WriteProfileData(profileData); err != nil {
		return err
	}

	return nil
}

func (repo ProfileRepository) ToggleIsMostUsed(skillName string, typeName string) error {
	profileData, err := repo.ReadProfileData()
	if err != nil {
		return err
	}

	isExists := false
	for i, skillType := range profileData.SkillSet{
		if skillType.TypeName == typeName {
			for k, skillItem := range profileData.SkillSet[i].SkillItems {
				if skillItem.Name == skillName{
					profileData.SkillSet[i].SkillItems[k].IsMostUsed = !profileData.SkillSet[i].SkillItems[k].IsMostUsed
					isExists = true
				}
			}
		}
	}

	if !isExists {
		return models.ErrSkillNotFound
	}

	// update the profile data
	if err := repo.WriteProfileData(profileData); err != nil {
		return err
	}

	return nil
}