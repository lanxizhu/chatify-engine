package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
