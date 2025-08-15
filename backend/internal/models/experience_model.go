package models

import (
	"time"

	"github.com/lib/pq"
)

type Experience struct {
	Id int `json:"id"`
	JobTitle string `json:"job_title"`
	CompanyName string `json:"company_name"`
	WorkPlace string `json:"work_place"`
	StartDate time.Time `json:"start_date"`
	EndDate pq.NullTime `json:"end_date"`
	Description string `json:"description"`
	Accomplishments []string `json:"accomplishments"`
	TechStack []string `json:"tech_stack"`
}