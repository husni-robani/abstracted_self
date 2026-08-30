package requests

import "mime/multipart"

type UpdateBlogRequest struct {
	Title           string                `form:"title" validate:"required"`
	Slug            string                `form:"slug"`
	Content         string                `form:"content" validate:"required"`
	BlogSnippet     string                `form:"blog_snippet"`
	Published       bool                  `form:"published"`
	ContentImageIds []int                 `form:"content_image_ids"`
	ImageFile       *multipart.FileHeader
}
