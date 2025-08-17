package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/auth"
	"github.com/husni-robani/abstracted_self/backend/internal/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	// set "multipart forms" data like image
	handler := handlers.SetupHandler(db)
	
	router := gin.Default()
	router.Use(auth.DefaultCorsConfig())

	// public router
	router.POST("/auth/login", handler.AuthHandler.Login)

	router.GET("/blogs", handler.BlogHandler.GetBlogs)
	router.GET("/blogs/:id", handler.BlogHandler.GetBlogByID)

	router.GET("/tech_types", handler.TechnologyTypeHandler.GetTypes)
	router.GET("/tech_types_technologies", handler.TechnologyTypeHandler.GetTypesWithTechnologies)

	router.GET("/technologies", handler.TechnologyHandler.GetTechnologies)

	router.GET("/experiences", handler.ExperienceHandler.GetExperiences)
	router.GET("/experiences/:id", handler.ExperienceHandler.GetExperienceById)

	router.GET("/projects", handler.ProjectHandler.GetProjects)
	router.GET("/projects/:id", handler.ProjectHandler.GetProjectById)

	router.GET("/assets/images/:file_name", handler.StorageHandler.GetImages)
	router.GET("/assets/documents/:file_name", handler.StorageHandler.GetDocuments)

	router.GET("/profile", handler.ProfileHandler.GetProfileData)


	// authorized router
	authorizedRouter := router.Group("", auth.AuthMiddleware())
	{
		authorizedRouter.POST("/auth/logout", handler.AuthHandler.Logout)
		authorizedRouter.POST("/auth/check", handler.AuthHandler.Check)
		authorizedRouter.GET("/auth/renew", handler.AuthHandler.RenewToken)
		
		authorizedRouter.POST("/blogs", handler.BlogHandler.CreateBlog)
		authorizedRouter.DELETE("/blogs/:id", handler.BlogHandler.DeleteBlog)
		authorizedRouter.PATCH("/blogs/:id", handler.BlogHandler.UpdateBlog)

		authorizedRouter.POST("/tech_types", handler.TechnologyTypeHandler.CreateTypes)
		authorizedRouter.DELETE("/tech_types/:id", handler.TechnologyTypeHandler.DeleteTypeByID)

		authorizedRouter.POST("/technologies", handler.TechnologyHandler.CreateTechnologies)
		authorizedRouter.DELETE("/technologies/:id", handler.TechnologyHandler.DeleteTechnology)

		authorizedRouter.POST("/experiences", handler.ExperienceHandler.CreateExperience)
		authorizedRouter.PATCH("/experiences/:id", handler.ExperienceHandler.UpdateExperience)
		authorizedRouter.DELETE("/experiences/:id", handler.ExperienceHandler.DeleteExperience)

		authorizedRouter.POST("/projects", handler.ProjectHandler.CreateProject)
		authorizedRouter.DELETE("/projects/:id", handler.ProjectHandler.DeleteProject)
		authorizedRouter.PUT("projects/:id", handler.ProjectHandler.UpdateProject)

		authorizedRouter.PUT("/profile", handler.ProfileHandler.UpdateProfile)
		authorizedRouter.POST("/profile/skill_type", handler.ProfileHandler.AddSkillType)
	}

	return router
}