package models

import "time"

type DailyVisit struct {
	ID int `json:"-"`
	UUID string `json:"uuid"`
	VisitDate time.Time `json:"visit_date"`
	Ip string `json:"ip"`
	Device string `json:"device"`
}

type DailyVisitCount struct {
	VisitDate string `json:"visit_date"`
	Count     int    `json:"count"`
}