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

type TypeWithTechnologies struct {
	TypeId int
	TypeName string
	TechID sql.NullInt64
	TechName sql.NullString
}

func (repo TechnologyTypeRepository) GetTypesWithTechnologies() ([]models.TechnologyType, error) {
	query := "SELECT tt.id, tt.type_name, t.id, t.name FROM technology_types tt LEFT JOIN technologies t ON tt.id = t.type_id"
	var typeTechs []TypeWithTechnologies

	rows, err := repo.db.Query(query)
	if err != nil {
		logger.Error.Printf("faield select query: %#v", err)
		return nil, err
	}

	for rows.Next(){
		var typeTech TypeWithTechnologies
		err := rows.Scan(&typeTech.TypeId, &typeTech.TypeName, &typeTech.TechID, &typeTech.TechName)
		if err != nil {
			logger.Error.Printf("failed scan row: %v", err)
			return nil, err
		}

		typeTechs = append(typeTechs, typeTech)
	}

	var technologyTypes []models.TechnologyType
	for i, v := range typeTechs {
		var technologyType models.TechnologyType
		if i == 0 {
			if v.TechID.Valid {
				technologyType = models.TechnologyType{
					Id: v.TypeId,
					TypeName: v.TypeName,
					Technologies: []models.Technology{
						{
							Id: int(v.TechID.Int64),
							Name: v.TechName.String,
						},
					},
				}
			}else {
				technologyType = models.TechnologyType{
					Id: v.TypeId,
					TypeName: v.TypeName,
					Technologies: nil,
				}
			}

			technologyTypes = append(technologyTypes, technologyType)
			continue
		}

		if technologyTypes[len(technologyTypes) - 1].Id != v.TypeId {
			technologyType = models.TechnologyType{
				Id: v.TypeId,
				TypeName: v.TypeName,
				Technologies: []models.Technology{
					{
						Id: int(v.TechID.Int64),
						Name: v.TechName.String,
					},
				},
			}

			technologyTypes = append(technologyTypes, technologyType)
			continue
		}

		technologyTypes[len(technologyTypes) - 1].Technologies = append(technologyTypes[len(technologyTypes)-1].Technologies, models.Technology{
			Id: int(v.TechID.Int64),
			Name: v.TechName.String,
		})

	}
	
	return technologyTypes, nil
}