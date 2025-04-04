package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          string     `json:"id"`
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
