package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/auth"
	"github.com/husni-robani/abstracted_self/backend/internal/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	// set "multipart forms" data like image
	handler := handlers.SetupHandler(db)

	router.POST("/auth/login", handler.AuthHandler.Login)

	router.GET("/blogs", handler.BlogHandler.GetBlogs)
	router.GET("/blogs/:id", handler.BlogHandler.GetBlogByID)

	router.GET("/tech_types", handler.TechnologyTypeHandler.GetTypes)
	router.GET("/tech_types_technologies", handler.TechnologyTypeHandler.GetTypesWithTechnologies)

	router.GET("/technologies", handler.TechnologyHandler.GetTechnologies)

	router.GET("/experiences", handler.ExperienceHandler.GetExperiences)
	router.GET("/experiences/:id", handler.ExperienceHandler.GetExperienceById)

	authorizedRouter := router.Group("")
	authorizedRouter.Use(auth.AuthMiddleware())
	{
		authorizedRouter.POST("/auth/logout", handler.AuthHandler.Logout)
		
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
	}

	return router
}