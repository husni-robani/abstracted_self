package requests

import "mime/multipart"

type UpdateBlogRequest struct {
	Title *string `form:"title,omitempty"`
	URL *string `form:"url,omitempty"`
	ImageFile *multipart.FileHeader `form:"file,omitempty"`
}