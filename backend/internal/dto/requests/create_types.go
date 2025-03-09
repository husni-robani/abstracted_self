package requests

type CreateTypesRequest struct {
	TypeNames []string `json:"type_names" validate:"required"`
}