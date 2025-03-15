package models

import "time"

type Experience struct {
	Id int `json:"id"`
	JobTitle string `json:"job_title"`
	CompanyName string `json:"company_name"`
	WorkPlace string `json:"work_place"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Description string `json:"description"`
	Accomplishments []string `json:"accomplishments"`
}