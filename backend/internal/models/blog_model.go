package models

type Blog struct {
	Id int `json:"id" validate:"numeric"`
	Title string `json:"title" form:"title" validate:"required"`
	URL string	`json:"url" form:"url" validate:"required,url"`
	Image string `json:"image" form:"image"`
	BlogSnippet string `json:"blog_snippet" form:"blog_snippet"`
}