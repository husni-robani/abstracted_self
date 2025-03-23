package models

import (
	"time"
)

type Project struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	TechStack []string `json:"tech_stack,omitempty"`
	SourceURL []string `json:"source_url,omitempty"`
	ProjectURL string `json:"project_url,omitempty"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Images []ProjectImage `json:"images,omitempty"`
}