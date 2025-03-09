package repositories

import (
	"database/sql"
	"fmt"

	"github.com/husni-robani/abstracted_self/backend/internal/logger"
	"github.com/husni-robani/abstracted_self/backend/internal/models"
)

type TechnologyTypeRepository struct {
	db *sql.DB
}

func NewTechnologyTypeRepository(db *sql.DB) TechnologyTypeRepository {
	return TechnologyTypeRepository{db: db}
}

func (repo TechnologyTypeRepository) GetAllType() ([]models.TechnologyType, error) {
	var types []models.TechnologyType

	rows, err := repo.db.Query("SELECT * FROM technology_types")
	if err != nil {
		logger.Error.Printf("failed to select technology types: %v", err)
		return nil, err
	}

	for rows.Next(){
		var techType models.TechnologyType

		err := rows.Scan(&techType.Id, &techType.TypeName)
		if err != nil {
			logger.Error.Printf("failed to scan row: %v", err)
			return nil, err
		}

		types = append(types, techType)
	}

	return types, nil
}

func (repo TechnologyTypeRepository) CreateTypes(typeNames []string) error {
	queryString := "INSERT INTO technology_types (type_name) values "

	for i, typeName := range typeNames {
		if i == len(typeNames) - 1 {
			queryString += fmt.Sprintf("('%s');", typeName)
		}else {
			queryString += fmt.Sprintf("('%s'), ", typeName)
		}
	}

	logger.Info.Printf("Query: %v\n", queryString)

	result, err := repo.db.Exec(queryString)
	if err != nil {
		logger.Error.Printf("failed to create technology type: %v", err)
		return err
	}

	totalRowAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed to get row affected");
		return err
	}

	logger.Info.Printf("Rows affected: %v", totalRowAffected)

	return nil
}

func (repo TechnologyTypeRepository) DeleteType(id int) error {
	result, err := repo.db.Exec("DELETE FROM technology_types WHERE id = $1", id)
	if err != nil {
		logger.Error.Printf("failed exec query delete: %v", err)
		return err
	}

	totalRowAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error.Printf("failed get total rows affected from query delete: %v", err)
		return err
	}

	logger.Info.Printf("Rows affected: %v", totalRowAffected)
	return nil
}