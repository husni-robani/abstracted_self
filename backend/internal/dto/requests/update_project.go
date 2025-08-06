package requests

import "mime/multipart"

type UpdateProjectRequest struct {
	Name string `form:"name,omitempty"`
	Description string `form:"description"`
	TechStack []string `form:"tech_stack"`
	SourceURL []string `form:"source_url"`
	ProjectURL string `form:"project_url"`
	StartDate string `form:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate string `form:"end_date" validate:"required,datetime=2006-01-02"`
	NewImages []multipart.FileHeader `form:"new_images"`
	ImagesToDelete []int `form:"images_to_delete"`
}