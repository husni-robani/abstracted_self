package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/dto/requests"
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

func (service DailyVisitService) ProfileVisitor(visitReq requests.VisitRequest, c *gin.Context) error {
	visit_data := models.DailyVisit{
		UUID: visitReq.UUID,
		VisitDate: time.Now(),
		Ip: c.ClientIP(),
		Device: c.GetHeader("User-Agent"),
	}

	if err := service.repo.StoreDailyVisit(visit_data); err != nil {
		return err
	}

	return nil
}