package models

import "time"

type Blog struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Image       string     `json:"image"`
	Content     string     `json:"content"`
	BlogSnippet string     `json:"blog_snippet"`
	Published   bool       `json:"published"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
