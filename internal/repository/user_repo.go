package repository

import (
	"chatify-engine/internal/model"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"time"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindUserByUsername(username string) (*model.User, error) {
	query := "SELECT * FROM user WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Nickname, &user.CreatedTime, &user.UpdatedTime, &user.LastTime)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO user (id, username, password, nickname, created_time, updated_time) VALUES (?, ?, ?, ?, ?, ?) "
	_, err := r.db.Exec(query, uuid.New().String(), user.Username, user.Password, user.Nickname, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateLoginTime(user *model.User) error {
	query := "UPDATE user SET last_time = ? WHERE id = ?"
	result, err := r.db.Exec(query, time.Now(), user.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
