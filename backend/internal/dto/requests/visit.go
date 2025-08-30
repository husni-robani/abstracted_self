package requests

type VisitRequest struct {
	UUID string `json:"uuid" validate:"required"`
	VisitDate string `json:"visit_date" validate:"required,datetime=2006-01-02"`
	Ip string `json:"ip"`
	Device string `json:"device"`
}