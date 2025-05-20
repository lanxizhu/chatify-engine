package repository

import (
	"chatify-engine/internal/model"
	"database/sql"
	"errors"
	"time"
)

type FriendRepository struct {
	db *sql.DB
}

func NewFriendRepository(db *sql.DB) *FriendRepository {
	return &FriendRepository{db}
}

func (r *FriendRepository) FindFriendsByUserID(userID string) ([]*model.Friend, error) {
	query := "SELECT friend, remark FROM friends WHERE mine = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	usersMap := make(map[string]*model.SimpleUser)
	userRepo := NewUserRepository(r.db)

	if simpleUsers, err := userRepo.FindAll(); err != nil {
		return nil, err
	} else {
		for _, user := range simpleUsers {
			usersMap[user.ID] = user
		}
	}

	var friends []*model.Friend
	for rows.Next() {
		var friend model.Friend
		if err = rows.Scan(&friend.ID, &friend.Remark); err != nil {
			return nil, err
		}

		friends = append(friends, &model.Friend{
			ID:       friend.ID,
			Account:  usersMap[friend.ID].Account,
			Username: usersMap[friend.ID].Username,
			Nickname: usersMap[friend.ID].Nickname,
			Avatar:   usersMap[friend.ID].Avatar,
			Remark:   friend.Remark,
		})
	}

	return friends, nil
}

func (r *FriendRepository) CheckFriendRequest(request *model.RequestFriend) (bool, error) {
	query := "SELECT CASE WHEN COUNT(*) > 0 THEN 'true' ELSE 'false' END FROM friend_request WHERE user = ? AND friend = ?"
	row := r.db.QueryRow(query, request.UserID, request.FriendID)

	var isExists bool
	err := row.Scan(&isExists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return isExists, nil
		}
		return isExists, err
	}

	return isExists, nil
}

func (r *FriendRepository) InsertFriendRequest(request *model.RequestFriend) error {
	query := "INSERT INTO friend_request (user, friend, remark) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, request.UserID, request.FriendID, &request.Remark)
	if err != nil {
		return err
	}

	return nil
}

func (r *FriendRepository) CheckFriendRequestByID(id string) (bool, error) {
	query := "SELECT CASE WHEN COUNT(*) > 0 THEN 'true' ELSE 'false' END FROM friend_request WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var isExists bool
	err := row.Scan(&isExists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return isExists, nil
		}
		return isExists, err
	}

	return isExists, nil
}

func (r *FriendRepository) UpdateFriendRequest(id string, status string) error {
	query := "UPDATE friend_request SET status = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, status, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
