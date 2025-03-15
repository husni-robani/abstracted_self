package repositories

import "database/sql"

type ExperienceRepo struct {
	db *sql.DB
}

func NewExperienceRepository(db *sql.DB) ExperienceRepo {
	return ExperienceRepo{
		db: db,
	}
}