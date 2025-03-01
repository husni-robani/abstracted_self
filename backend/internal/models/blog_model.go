package models

type Blog struct {
	Id int `json:"id"`
	Title string `json:"title"`
	URL string	`json:"url"`
	Image string `json:"image"`
}