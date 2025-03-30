package router

import (
	"chatify-engine/internal/handler"
	"chatify-engine/internal/repository"
	"chatify-engine/internal/service"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(db *sql.DB) *gin.Engine {
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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	publicGroup := router.Group("/api/v1")
	{
		publicGroup.POST("/register", userHandler.Register)
	}

	return router
}
