package services

import (
	"math"
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

func (service DailyVisitService) GetDailyVisitCounts(startDate, endDate string) ([]models.DailyVisitCount, error) {
	counts, err := service.repo.GetDailyVisitCounts(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return counts, nil
}

func (service DailyVisitService) GetDailyVisits(startDate, endDate string, page, limit int) ([]models.DailyVisit, models.Pagination, error) {
	total, err := service.repo.CountDailyVisits(startDate, endDate)
	if err != nil {
		return nil, models.Pagination{}, err
	}

	offset := (page - 1) * limit
	visits, err := service.repo.GetDailyVisits(startDate, endDate, limit, offset)
	if err != nil {
		return nil, models.Pagination{}, err
	}

	pagination := models.Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
	}

	return visits, pagination, nil
}