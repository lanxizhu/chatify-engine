package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type MediaMode string

const (
	AvatarMode MediaMode = "avatar"
)

func GetMediaUrl(c *gin.Context, mode MediaMode) string {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}

	host := c.Request.Host
	return fmt.Sprintf("%s://%s/media/%s", scheme, host, mode)
}
