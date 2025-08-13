package requests

type UpdateExperienceRequest struct {
	JobTitle *string `json:"job_title,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	WorkPlace *string `json:"work_place,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	Description *string `json:"description,omitempty"`
	Accomplishments []string `json:"accomplishments,omitempty"`
	TechStack []string `json:"tech_stack,omitempty"`
}