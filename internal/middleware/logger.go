package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mssola/useragent"
	"go.uber.org/zap"
)

var globalLogger *zap.Logger

func init() {
	var err error

	if gin.Mode() == gin.ReleaseMode {
		globalLogger, err = zap.NewProduction()
	} else {
		globalLogger, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

// LoggerMiddleware is a Gin middleware that creates a logger instance with request_id for each request
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// create a new logger instance for each request
		requestID := uuid.New().String()
		ua := useragent.New(c.Request.UserAgent())
		browser, _ := ua.Browser()
		requestLogger := globalLogger.With(
			zap.String("request_id", requestID),
			zap.String("client_ip", c.ClientIP()),
			zap.String("Browser", browser),
			zap.String("Os", ua.OS()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		// store the logger in the Gin Context
		c.Set("logger", requestLogger)

		// Process the next request
		c.Next()
	}
}

// GetLoggerFromContext retrieves the logger from the Gin Context
func GetLoggerFromContext(c *gin.Context) *zap.Logger {
	if logger, exists := c.Get("logger"); exists {
		if zapLogger, ok := logger.(*zap.Logger); ok {
			return zapLogger
		}
	}
	// If the logger is not found in the context, return the global logger
	return globalLogger.With(zap.String("context", "logger_not_found"))
}
