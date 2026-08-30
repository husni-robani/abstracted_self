package models

import "time"

type Blog struct {
	Id           int        `json:"id"`
	Title        string     `json:"title"`
	Slug         string     `json:"slug"`
	Content      string     `json:"content,omitempty"`
	BlogSnippet  string     `json:"blog_snippet"`
	Published    bool       `json:"published"`
	CoverImageId int        `json:"-"`
	CoverImage   *Image     `json:"cover_image"`
	Images       []Image    `json:"images,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
