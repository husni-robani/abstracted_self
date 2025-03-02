package models

type Blog struct {
	Id int `json:"id,omitempty" validate:"numeric"`
	Title string `json:"title" validate:"required"`
	URL string	`json:"url" validate:"required,url"`
	Image string `json:"image" validate:"required,filepath"`
}