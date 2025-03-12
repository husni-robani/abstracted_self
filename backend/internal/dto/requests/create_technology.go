package requests

type CreateTechnologyRequest struct {
	TypeId int `json:"type_id" validate:"required"`
	Name string `json:"name" validate:"required"`
}