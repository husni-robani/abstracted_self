package models

import "time"

type DailyVisit struct {
	ID int `json:"id"`	
	UUID string `json:"uuid"`
	VisitDate time.Time `json:"visit_date"`
	Ip string `json:"ip"`
	Device string `json:"device"`
}