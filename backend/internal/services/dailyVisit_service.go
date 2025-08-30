package services

import "github.com/husni-robani/abstracted_self/backend/internal/repositories"

type DailyVisitService struct {
	repo repositories.DailyVisitRepository
}

func NewDailyVisitService(repo repositories.DailyVisitRepository) DailyVisitService{
	return DailyVisitService{
		repo: repo,
	}
}