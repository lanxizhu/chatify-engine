package router

import (
	"chatify-engine/internal/handler"
	"chatify-engine/internal/middleware"
	"chatify-engine/internal/repository"
	"chatify-engine/internal/service"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Create(db *sql.DB) *gin.Engine {
	gin.ForceConsoleColor()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.Use(middleware.LoggerMiddleware())

	if mode := gin.Mode(); mode != gin.TestMode {
		router.LoadHTMLGlob("templates/*")
		router.Static("/media", "./uploads")
		router.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
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

		logger := middleware.GetLoggerFromContext(c)
		logger.Info("Ping request received")
	})

	wsHandler := handler.SetupWsHandler()
	router.Handle(http.MethodGet, "/websocket/:id", wsHandler.Connect)

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
		userGroup.GET("/search", userHandler.SearchUsers)
		userGroup.GET("/find", userHandler.FindUser)
	}

	friendRepo := repository.NewFriendRepository(db)
	friendService := service.NewFriendService(friendRepo)
	friendHandler := handler.NewFriendHandler(friendService)

	friendGroup := protectedGroup.Group("/friend")
	{
		friendGroup.GET("/", friendHandler.GetFriends)
	}

	return router
}
