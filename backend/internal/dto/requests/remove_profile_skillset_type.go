package requests

type RemoveProfileSkillSetType struct {
	TypeName string `form:"type_name" validate:"required"`
}