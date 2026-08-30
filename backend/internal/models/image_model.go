package models

import (
	"errors"
	"fmt"
)

var ErrImageNotFound = errors.New("image not found")

type Image struct {
	Id       int    `json:"id"`
	URL      string `json:"url"`
	FileName string `json:"file_name"`
	FileSize int    `json:"file_size"`
	MimeType string `json:"mime_type"`
}

func ImageURL(id int) string {
	return fmt.Sprintf("/images/%d", id)
}
