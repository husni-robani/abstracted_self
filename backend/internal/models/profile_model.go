package models

import "errors"


var ErrSkillTypeDuplicate = errors.New("skill type already exists")
var ErrSkillDuplicate = errors.New("skill already exists")
var ErrSkillTypeNotFound = errors.New("type not found")

type Profile struct {
	Name string `json:"name"`
	Summary string `json:"summary"`
	Bio string `json:"bio"`
	Taglines []string `json:"taglines"`
	ResumeFileName string `json:"resume_file_name"`
	SkillSet []SkillType `json:"skill_set"`
}

type SkillType struct {
	TypeName string `json:"type_name"`
	SkillItems []Skill `json:"skill_items"`
}

type Skill struct {
	Name string `json:"name"`
	IconFilename string `json:"icon_filename"`
	IsMostUsed bool `json:"is_most_used"`
}

