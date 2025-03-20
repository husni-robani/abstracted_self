package requests

type CreateExperienceRequest struct {
	JobTitle string `json:"job_title" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	WorkPlace string `json:"work_place" validate:"required"`
	StartDate string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate string `json:"end_date" validate:"required,datetime=2006-01-02"`
	Description string `json:"description"`
	Accomplishments []string `json:"accomplishments"`
}