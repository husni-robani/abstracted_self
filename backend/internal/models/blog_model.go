package models

import "mime/multipart"

type Blog struct {
	Id int `json:"id,omitempty" validate:"numeric"`
	Title string `json:"title" form:"title" validate:"required"`
	URL string	`json:"url" form:"url" validate:"required,url"`
	Image string `json:"image,omitempty" form:"image,omitempty"`
	ImageFile *multipart.FileHeader
}