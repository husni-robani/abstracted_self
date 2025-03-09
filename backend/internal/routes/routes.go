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

	authorizedRouter := router.Group("")
	authorizedRouter.Use(auth.AuthMiddleware())
	{
		authorizedRouter.POST("/auth/logout", handler.AuthHandler.Logout)
		
		authorizedRouter.POST("/blogs", handler.BlogHandler.CreateBlog)
		authorizedRouter.DELETE("/blogs/:id", handler.BlogHandler.DeleteBlog)
		authorizedRouter.PATCH("/blogs/:id", handler.BlogHandler.UpdateBlog)
	}

	return router
}