package requests

type UpdateExperienceRequest struct {
	JobTitle *string `json:"job_title,omitempty"`
	CompanyName *string `json:"company_name,omitempty"`
	WorkPlace *string `json:"work_place,omitempty"`
	StartDate *string `json:"start_date,omitempty" validate:"datetime=2006-01-02"`
	EndDate *string `json:"end_date,omitempty" validate:"datetime=2006-01-02"`
	Description *string `json:"description,omitempty"`
	Accomplishments []string `json:"accomplishments,omitempty"`
}