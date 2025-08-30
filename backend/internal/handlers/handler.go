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
	TechnologyHandler *TechnologyHandler
	ExperienceHandler *ExperienceHandler
	ProjectHandler *ProjectHandler
	StorageHandler *StorageHandler
	ProfileHandler *ProfileHandler
	DailyVisitHandler *DailyVisitHandler
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

	technologyRepo := repositories.NewTechnologyRepository(db)
	technologyService := services.NewTechnologyService(technologyRepo)
	technologyHandler := NewTechnologyHandler(technologyService)

	experienceRepo := repositories.NewExperienceRepository(db)
	experienceService := services.NewExperienceService(experienceRepo)
	experienceHandler := NewExperienceHandler(experienceService)

	projectRepo := repositories.NewProjectRepository(db)
	projectImageRepo := repositories.NewProjectImageRepository(db)
	projectService := services.NewProjectService(projectRepo, projectImageRepo)
	projectHandler := NewProjectHandler(projectService)

	storageHandler := NewStorageHandler()

	profileRepo := repositories.NewProfileRepository()
	profileService := services.NewProfileService(profileRepo)
	ProfileHandler := NewProfileHandler(profileService)

	dailyVisitRepo := repositories.NewDailyVisitRepository(db)
	dailyVisitService := services.NewDailyVisitService(dailyVisitRepo)
	dailyVisitHandler := NewDailyVisitHandler(dailyVisitService)

	return Handlers{
		BlogHandler: &blogHandler,
		AuthHandler: &authHandler,
		TechnologyTypeHandler: &techTypeHandler,
		TechnologyHandler: &technologyHandler,
		ExperienceHandler: &experienceHandler,
		ProjectHandler: &projectHandler,
		StorageHandler: &storageHandler,
		ProfileHandler: &ProfileHandler,
		DailyVisitHandler: &dailyVisitHandler,
	}
}