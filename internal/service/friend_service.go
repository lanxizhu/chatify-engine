package service

import (
	"chatify-engine/internal/model"
	"chatify-engine/internal/repository"
	"errors"
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

func (s *FriendService) RequestFriend(request *model.RequestFriend) error {
	isExists, err := s.friendRepo.CheckFriendRequest(request)
	if err != nil {
		return err
	}
	if isExists {
		return errors.New("this friend is already requested")
	}
	if err = s.friendRepo.InsertFriendRequest(request); err != nil {
		return err
	}
	return nil
}
