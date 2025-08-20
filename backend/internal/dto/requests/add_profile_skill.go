package requests

import "mime/multipart"

type AddProfileSkill struct {
	Name string `form:"name" validate:"required"`
	IsMostUsed *bool `form:"is_most_used" validate:"required"`
	IconFile *multipart.FileHeader
	SkillType string `form:"skill_type" validate:"required"`
	URL string `form:"url" validate:"required"`
}