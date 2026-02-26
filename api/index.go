package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirna70/api-quest-test/routes"
)

var app *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	app = gin.New()
	app.Use(gin.Recovery())
	routes.SetupRoutes(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
