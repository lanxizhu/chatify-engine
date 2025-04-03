package router

import (
	"chatify-engine/internal/handler"
	"chatify-engine/internal/middleware"
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
		router.Static("/media", "./uploads")
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
		publicGroup.POST("/login", userHandler.Login)
	}

	protectedGroup := router.Group("/api/v1")
	protectedGroup.Use(middleware.AuthMiddleware())
	protectedGroup.GET("/validateToken", userHandler.ValidateToken)

	userGroup := protectedGroup.Group("/user")
	{
		userGroup.PUT("/changePassword", userHandler.ChangePassword)
		userGroup.POST("/uploadAvatar", userHandler.UploadAvatar)
		userGroup.PUT("/updateInfo", userHandler.UpdateUser)
	}

	return router
}
