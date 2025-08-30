package handlers

import (
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type DailyVisitHandler struct {
	service services.DailyVisitService
}

func NewDailyVisitHandler(service services.DailyVisitService) DailyVisitHandler {
	return DailyVisitHandler{
		service: service,
	}
}