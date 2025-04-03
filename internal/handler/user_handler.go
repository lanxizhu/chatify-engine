package handler

import (
	"chatify-engine/internal/model"
	"chatify-engine/internal/service"
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
		avatarURL := fmt.Sprintf("http://localhost:8888/media/avatars/%s", *user.Avatar)
		avatar = &avatarURL
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
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

	if err = c.SaveUploadedFile(file, "./uploads/avatars/"+filename); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"url":     "http://localhost:8888/media/avatars/" + filename,
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

	avatarURL := fmt.Sprintf("http://localhost:8888/media/avatars/%s", *user.Avatar)
	user.Avatar = &avatarURL

	c.JSON(http.StatusOK, gin.H{
		"message": "Update success",
		"info":    user,
	})
}
