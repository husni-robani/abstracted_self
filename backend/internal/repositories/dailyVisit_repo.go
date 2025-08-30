package repositories

import "database/sql"

type DailyVisitRepository struct {
	DB *sql.DB
}

func NewDailyVisitRepository(db *sql.DB) DailyVisitRepository{
	return DailyVisitRepository{
		DB: db,
	}
}