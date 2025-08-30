package requests

type VisitRequest struct {
	UUID string `json:"uuid" validate:"required"`
}