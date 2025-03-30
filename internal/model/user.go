package model

import (
	"time"
)

type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	Nickname    *string   `json:"nickname"`
	Avatar      *string   `json:"avatar"`
	CreatedTime time.Time `json:"created_at"`
	UpdatedTime time.Time `json:"updated_at"`
	LastTime    time.Time `json:"last_login"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
