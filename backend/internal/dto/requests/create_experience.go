package requests

type CreateExperienceRequest struct {
	JobTitle string `json:"job_title"`
	CompanyName string `json:"company_name"`
	WorkPlace string `json:"work_place"`
	StartDate string `json:"start_date"`
}