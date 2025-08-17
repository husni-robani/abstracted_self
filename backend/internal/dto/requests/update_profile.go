package requests

import "mime/multipart"

type UpdateProfileRequest struct {
	Name string `form:"name,omitempty"`
	Summary string `form:"summary,omitempty"`
	Bio string `form:"bio,omitempty"`
	Taglines []string `form:"taglines,omitempty"`
	ResumeFile *multipart.FileHeader
}