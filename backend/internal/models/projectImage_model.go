package models

type ProjectImage struct {
	Id int `json:"id,omitempty"`
	ProjectId int `json:"project_id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	FileSize int `json:"file_size,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
}