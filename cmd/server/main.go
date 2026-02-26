package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirna70/api-quest-test/internal/routes"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
