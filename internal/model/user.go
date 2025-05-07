package model

import (
	"chatify-engine/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type User struct {
	ID          string     `json:"id"`
	Account     uint       `json:"account"`
	Username    string     `json:"username"`
	Password    string     `json:"-"`
	Nickname    *string    `json:"nickname"`
	Avatar      *string    `json:"avatar"`
	CreatedTime time.Time  `json:"created_at"`
	UpdatedTime time.Time  `json:"updated_at"`
	LastTime    *time.Time `json:"last_login"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UploadAvatar struct {
	ID   string `json:"id"`
	File string `json:"file"`
}

type SimpleUser struct {
	ID       string     `json:"id"`
	Account  uint       `json:"account"`
	Username string     `json:"username"`
	Nickname *string    `json:"nickname"`
	Avatar   *string    `json:"avatar"`
	LastTime *time.Time `json:"last_login"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateAccount() {
	u.Account = 10000000 + uint(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(10000000))
}

func (u *SimpleUser) GetAvatarUrl(c *gin.Context) {
	if u.Avatar == nil {
		return
	}

	avatarURL := fmt.Sprintf("%s/%s", utils.GetMediaUrl(c, utils.AvatarMode), *u.Avatar)
	u.Avatar = &avatarURL
}
