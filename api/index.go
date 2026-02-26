package main

import (
	"net/http"

	"github.com/sirna70/api-quest-test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	routes.SetupRoutes(r)

	http.ListenAndServe(":8080", r)
}
