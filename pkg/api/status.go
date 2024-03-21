package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "HTTP REST Server up and running",
	})

}
