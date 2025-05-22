package handler

import (
	"chatify-engine/internal/middleware"
	"chatify-engine/internal/model"
	"chatify-engine/internal/service"
	"chatify-engine/pkg/redis"
	"chatify-engine/pkg/utils"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request model.RegisterUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.userService.Register(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "register success",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var request model.LoginUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	token, user, err := h.userService.Login(request.Username, request.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	var avatar *string
	if user.Avatar != nil {
		avatarURL := fmt.Sprintf("%s/%s", utils.GetMediaUrl(c, utils.AvatarMode), *user.Avatar)
		avatar = &avatarURL
	}

	logger := middleware.GetLoggerFromContext(c)
	logger.Info("User login")

	rdb := redis.GetRdb()
	rdb.Set(c, "token:"+user.ID, token, 0)

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"account":    user.Account,
		"username":   user.Username,
		"nickname":   user.Nickname,
		"avatar":     &avatar,
		"created_at": user.CreatedTime,
		"updated_at": user.UpdatedTime,
		"last_login": user.LastTime,
		"token":      token,
	})
}

func (h *UserHandler) ValidateToken(c *gin.Context) {
	ID, _ := c.Get("user_id")
	Username, _ := c.Get("username")
	c.Get("nickname")
	c.JSON(http.StatusOK, gin.H{
		"id":       ID,
		"username": Username,
	})
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	var request model.ChangePassword

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	id, _ := c.Get("user_id")
	err := h.userService.ChangePassword(id.(string), request.OldPassword, request.NewPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
	})
}

func (h *UserHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "No file uploaded",
		})
		return
	}

	id, _ := c.Get("user_id")

	shaId := sha256.Sum256([]byte(id.(string)))

	filetype := strings.Split(file.Filename, ".")[1]
	filename := fmt.Sprintf("%x.%s", shaId, filetype)

	if err = c.SaveUploadedFile(file, fmt.Sprintf("./uploads/%s/%s", utils.AvatarMode, filename)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save file",
		})
		return
	}

	if err = h.userService.UploadAvatar(id.(string), filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	avatar := fmt.Sprintf("%s/%s", utils.GetMediaUrl(c, utils.AvatarMode), filename)
	c.JSON(http.StatusOK, gin.H{
		"url":     avatar,
		"message": "Upload success",
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var request model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	id, _ := c.Get("user_id")
	request.ID = id.(string)

	user, err := h.userService.UpdateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	avatarURL := fmt.Sprintf("%s/%s", utils.GetMediaUrl(c, utils.AvatarMode), *user.Avatar)
	user.Avatar = &avatarURL

	c.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"info":    user,
	})
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	name := c.Query("name")
	users, err := h.userService.SearchUsers(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if users == nil {
		c.JSON(http.StatusOK, gin.H{
			"users":   []interface{}{},
			"message": "No users found",
		})
		return
	}

	for _, user := range users {
		user.GetAvatarUrl(c)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) FindUser(c *gin.Context) {
	user, err := h.userService.FindUser(c.Query("keyword"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusOK, gin.H{
			"user":    nil,
			"message": "No user found",
		})
		return
	}

	user.GetAvatarUrl(c)
	c.JSON(http.StatusOK, user)
}
