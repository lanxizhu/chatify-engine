package repository

import (
	"chatify-engine/internal/model"
	"database/sql"
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
