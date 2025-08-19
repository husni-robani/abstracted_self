package requests

type ToggleSkillMostUsed struct {
	SkillName string `form:"skill_name" validate:"required"`
	TypeName string `form:"type_name" validate:"required"`
}