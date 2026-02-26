package routes

import (
	"api-quest-test/internal/handlers"
	"api-quest-test/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
	r.POST("/echo", handlers.EchoHandler)

	r.POST("/auth/token", handlers.GenerateToken)

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("", handlers.CreateBook)
		bookRoutes.GET("", middleware.AuthMiddleware(), handlers.GetBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}
}
