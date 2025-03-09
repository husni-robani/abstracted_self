package repositories

import "database/sql"

type TechnologyTypeRepository struct {
	db *sql.DB
}

func NewTechnologyTypeRepository(db *sql.DB) TechnologyTypeRepository {
	return TechnologyTypeRepository{db: db}
}

