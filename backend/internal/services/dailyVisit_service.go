package services

import (
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
)

type DailyVisitService struct {
	repo repositories.DailyVisitRepository
}

func NewDailyVisitService(repo repositories.DailyVisitRepository) DailyVisitService{
	return DailyVisitService{
		repo: repo,
	}
}

func (service DailyVisitService) ProfileVisitor(visitReq requests.VisitRequest) error {
	dateFormatted, err := time.Parse("2006-01-02", visitReq.VisitDate)
	if err != nil {
		logger.Error.Printf("Failed to parse VisitDate: %v", err)
		return err
	}

	visit_data := models.DailyVisit{
		UUID: visitReq.UUID,
		VisitDate: dateFormatted,
		Ip: visitReq.Ip,
		Device: visitReq.Device,
	}

	if err := service.repo.StoreDailyVisit(visit_data); err != nil {
		return err
	}

	return nil
}