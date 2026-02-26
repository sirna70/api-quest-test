package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirna70/api-quest-test/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
	r.POST("/echo", handlers.EchoHandler)

	r.POST("/auth/token", handlers.GenerateToken)

	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("", handlers.CreateBook)
		bookRoutes.GET("", handlers.GetBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}
}
