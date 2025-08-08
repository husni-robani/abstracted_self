package models

type ProjectImage struct {
	Id int `json:"id"`
	ProjectId int `json:"project_id"`
	FileName string `json:"file_name"`
	FileSize int `json:"file_size"`
	MimeType string `json:"mime_type"`
	ImageURL string `json:"image_url"`
}