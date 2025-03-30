package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create() *gin.Engine {
	gin.ForceConsoleColor()

	router := gin.Default()

	if mode := gin.Mode(); mode != gin.TestMode {
		router.LoadHTMLGlob("templates/*")
	}

	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Hello, World!",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
