package handlers

import (
	"database/sql"

	"github.com/husni-robani/abstracted_self/backend/internal/repositories"
	"github.com/husni-robani/abstracted_self/backend/internal/services"
)

type Handlers struct {
	BlogHandler *BlogHandler
	AuthHandler *AuthHandler
	TechnologyTypeHandler *TechnologyTypeHandler
}

func SetupHandler(db *sql.DB) Handlers {
	blogRepo := repositories.NewBlogRepository(db)
	blogService := services.NewBlogService(blogRepo)
	blogHandler := NewBlogHandler(blogService)

	authService := services.NewAuthService()
	authHandler := NewAuthHandler(authService)

	techTypeRepo := repositories.NewTechnologyTypeRepository(db)
	techTypeService := services.NewTechnologyTypeService(techTypeRepo)
	techTypeHandler := NewTechnologyTypeHandler(techTypeService)

	return Handlers{
		BlogHandler: &blogHandler,
		AuthHandler: &authHandler,
		TechnologyTypeHandler: &techTypeHandler,
	}
}