package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func EchoHandler(c *gin.Context) {
	var raw json.RawMessage

	if err := c.ShouldBindJSON(&raw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	c.Data(http.StatusOK, "application/json", raw)
}
