package models

import (
	"time"
)

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TechStack []string `json:"tech_stack"`
	SourceURL []string `json:"source_url"`
	ProjectURL string `json:"project_url"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Images []ProjectImage `json:"images"`
}