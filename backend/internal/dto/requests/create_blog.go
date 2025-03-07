package requests

import "mime/multipart"

type CreateBlogRequest struct{
	Title string `form:"title" validate:"required"`
	URL string	`form:"url" validate:"required,url"`
	Image string `form:"image,omitempty"`
	ImageFile *multipart.FileHeader
}