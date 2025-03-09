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
	router.POST("/blogs", handler.BlogHandler.CreateBlog)
	router.DELETE("/blogs/:id", handler.BlogHandler.DeleteBlog)
	router.PATCH("/blogs/:id", handler.BlogHandler.UpdateBlog)

	authorizedRouter := router.Group("")
	authorizedRouter.Use(auth.AuthMiddleware())
	{
		authorizedRouter.POST("/auth/logout", handler.AuthHandler.Logout)
	}

	return router
}