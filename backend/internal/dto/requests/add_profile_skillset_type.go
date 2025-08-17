package requests

type AddProfileSkillSetType struct {
	TypeName string `json:"type_name" validate:"required"`
}