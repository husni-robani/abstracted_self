package requests

type RemoveProfileSkill struct {
	SkillName string `form:"skill_name" validate:"required"`
}