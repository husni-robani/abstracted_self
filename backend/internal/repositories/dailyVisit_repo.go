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

func (repo DailyVisitRepository) GetDailyVisitCounts(startDate, endDate string) ([]models.DailyVisitCount, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `
		SELECT d.day::date AS visit_date, COUNT(dv.uuid)::int AS count
		FROM generate_series($1::date, $2::date, interval '1 day') AS d(day)
		LEFT JOIN daily_visits dv ON dv.visit_date = d.day
		GROUP BY d.day
		ORDER BY d.day`

	rows, err := repo.DB.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		logger.Error.Printf("failed exec query select daily visit counts: %v", err)
		return nil, err
	}
	defer rows.Close()

	var counts []models.DailyVisitCount

	for rows.Next() {
		var count models.DailyVisitCount
		var visitDate time.Time

		if err := rows.Scan(&visitDate, &count.Count); err != nil {
			logger.Error.Printf("failed to scan daily visit count: %v", err)
			return nil, err
		}

		count.VisitDate = visitDate.Format("2006-01-02")
		counts = append(counts, count)
	}

	return counts, nil
}

func (repo DailyVisitRepository) GetDailyVisits(startDate, endDate string, limit, offset int) ([]models.DailyVisit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `
		SELECT uuid, visit_date, ip, device
		FROM daily_visits
		WHERE visit_date >= $1::date AND visit_date <= $2::date
		ORDER BY visit_date DESC, id DESC
		LIMIT $3 OFFSET $4`

	rows, err := repo.DB.QueryContext(ctx, query, startDate, endDate, limit, offset)
	if err != nil {
		logger.Error.Printf("failed exec query select daily visits: %v", err)
		return nil, err
	}
	defer rows.Close()

	var visits []models.DailyVisit

	for rows.Next() {
		var visit models.DailyVisit
		var ip sql.NullString
		var device sql.NullString

		if err := rows.Scan(&visit.UUID, &visit.VisitDate, &ip, &device); err != nil {
			logger.Error.Printf("failed to scan daily visit: %v", err)
			return nil, err
		}

		visit.Ip = ip.String
		visit.Device = device.String
		visits = append(visits, visit)
	}

	return visits, nil
}

func (repo DailyVisitRepository) CountDailyVisits(startDate, endDate string) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var total int
	query := `
		SELECT COUNT(*)::int
		FROM daily_visits
		WHERE visit_date >= $1::date AND visit_date <= $2::date`

	if err := repo.DB.QueryRowContext(ctx, query, startDate, endDate).Scan(&total); err != nil {
		logger.Error.Printf("failed to count daily visits: %v", err)
		return 0, err
	}

	return total, nil
}