package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InfoEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "test",
	})
}
