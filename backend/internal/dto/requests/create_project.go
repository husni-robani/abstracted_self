package requests

import "mime/multipart"

type CreateProjectRequest struct {
	Name string `form:"name" validate:"required"`
	Description string `form:"description,omitempty"`
	TechStack []string `form:"tech_stack,omitempty"`
	SourceURL []string `form:"source_url,omitempty"`
	ProjectURL string `form:"project_url,omitempty"`
	StartDate string `form:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate string	`form:"end_date" validate:"required,datetime=2006-01-02"`
	Images []multipart.FileHeader `form:"images" validate:"required"`
}