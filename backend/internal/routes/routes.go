package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/husni-robani/abstracted_self/backend/internal/handlers"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	// set "multipart forms" data like image
	handler := handlers.SetupHandler(db)

	router.GET("/blogs", handler.BlogHandler.GetBlogs)
	router.GET("/blogs/:id", handler.BlogHandler.GetBlogByID)
	router.POST("/blogs", handler.BlogHandler.CreateBlog)
	router.DELETE("/blogs/:id", handler.BlogHandler.DeleteBlog)
	router.PATCH("/blogs/:id", handler.BlogHandler.UpdateBlog)

	return router
}