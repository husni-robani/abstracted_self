package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type DailyVisitRepository struct {
	DB *sql.DB
}

func NewDailyVisitRepository(db *sql.DB) DailyVisitRepository{
	return DailyVisitRepository{
		DB: db,
	}
}

func (repo DailyVisitRepository) StoreDailyVisit(data models.DailyVisit) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	query := "INSERT INTO daily_visits (uuid, visit_date, ip, device) VALUES ($1, $2, $3, $4)"
	result, err := repo.DB.ExecContext(ctx, query, &data.UUID, &data.VisitDate, &data.Ip, &data.Device)
	if err != nil {
		logger.Error.Printf("Insert daily visit failed: %v", err.Error())
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	logger.Info.Printf("Insert daily_visits successful | Rows Affected: %d", rowsAffected)

	return nil
}