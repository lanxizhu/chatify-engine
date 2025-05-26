package handler

import (
	"chatify-engine/internal/model"
	"chatify-engine/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FriendHandler struct {
	friendService *service.FriendService
}

func NewFriendHandler(friendService *service.FriendService) *FriendHandler {
	return &FriendHandler{friendService}
}

func (h *FriendHandler) GetFriends(c *gin.Context) {
	id, _ := c.Get("user_id")
	friends, err := h.friendService.FindFriends(id.(string))

	for _, friend := range friends {
		friend.GetAvatarUrl(c)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve friends",
		})
		return
	}

	if len(friends) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"friends": []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func (h *FriendHandler) AddFriend(c *gin.Context) {
	id, _ := c.Get("user_id")
	friendID := c.Param("friend_id")

	request := &model.RequestFriend{
		UserID:   id.(string),
		FriendID: friendID,
	}

	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	err := h.friendService.RequestFriend(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to request friend: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Friend request sent successfully",
	})
}

func (h *FriendHandler) GetFriendRequests(c *gin.Context) {
	id, _ := c.Get("user_id")
	requests, err := h.friendService.GetFriendRequests(id.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve friend requests: " + err.Error(),
		})
		return
	}

	if len(requests) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"requests": []interface{}{},
		})
		return
	}

	for _, request := range requests {
		request.GetAvatarUrl(c)
	}

	c.JSON(http.StatusOK, requests)
}

func (h *FriendHandler) HandleRequest(c *gin.Context) {
	id := c.Param("request_id")

	action := &model.HandleRequest{
		ID: id,
	}

	if err := c.ShouldBindJSON(action); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if action.Status == "" || action.Status != "accept" && action.Status != "reject" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request status",
		})
		return
	}

	if err := h.friendService.HandleRequest(action); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to handle request: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Request handled successfully",
	})
}
