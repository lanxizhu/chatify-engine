package service

import (
	"chatify-engine/internal/model"
	"chatify-engine/internal/repository"
)

type FriendService struct {
	friendRepo *repository.FriendRepository
}

func NewFriendService(friendRepo *repository.FriendRepository) *FriendService {
	return &FriendService{friendRepo}
}

func (s *FriendService) FindFriends(userID string) ([]*model.Friend, error) {
	friends, err := s.friendRepo.FindFriendsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return friends, nil
}
