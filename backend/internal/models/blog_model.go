package models

import "mime/multipart"

type Blog struct {
	Id int `json:"id" validate:"numeric"`
	Title string `json:"title" form:"title" validate:"required"`
	URL string	`json:"url" form:"url" validate:"required,url"`
	Image string `json:"image" form:"image"`
	ImageFile *multipart.FileHeader
}