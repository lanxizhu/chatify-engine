package model

import (
	"chatify-engine/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Friend struct {
	ID       string  `json:"id"`
	Account  uint    `json:"account"`
	Username string  `json:"username"`
	Nickname *string `json:"nickname"`
	Avatar   *string `json:"avatar"`
	Remark   *string `json:"remark"`
}

type RequestFriend struct {
	UserID   string  `json:"user_id"`
	FriendID string  `json:"friend_id"`
	Remark   *string `json:"remark"`
}

type FriendRequest struct {
	ID     string  `json:"id"`
	User   string  `json:"user"`
	Remark *string `json:"remark"`
	Status string  `json:"status"`
	Type   string  `json:"type"`
}

type HandleRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

func (f *Friend) GetAvatarUrl(c *gin.Context) {
	if f.Avatar == nil {
		return
	}

	avatarURL := fmt.Sprintf("%s/%s", utils.GetMediaUrl(c, utils.AvatarMode), *f.Avatar)
	f.Avatar = &avatarURL
}
