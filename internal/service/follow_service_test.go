package service

import (
	"errors"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/response"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository"
)

type FollowService interface {
	FollowUser(followerID, followedID uint) error
	UnfollowUser(followerID, followedID uint) error
	GetFollowersCount(userID uint) (*response.FollowersCountResponse, error)
	GetFollowersList(userID uint, order string) (*response.FollowersListResponse, error)
	GetFollowedList(userID uint, order string) (*response.FollowedListResponse, error)
}

type followServiceImpl struct {
	followRepo repository.FollowRepository
	userRepo   repository.UserRepository
}

func NewFollowService(followRepo repository.FollowRepository, userRepo repository.UserRepository) FollowService {
	return &followServiceImpl{
		followRepo: followRepo,
		userRepo:   userRepo,
	}
}


func (s *followServiceImpl) FollowUser(followerID, followedID uint) error {
	
	if followerID == followedID {
		return errors.New("cannot follow yourself")
	}

	
	if !s.userRepo.ExistsByID(followerID) {
		return errors.New("follower user not found")
	}
	if !s.userRepo.ExistsByID(followedID) {
		return errors.New("followed user not found")
	}

	
	if s.followRepo.IsFollowing(followerID, followedID) {
		return errors.New("already following this user")
	}

	follow := &domain.Follow{
		FollowerID: followerID,
		FollowedID: followedID,
	}

	return s.followRepo.Create(follow)
}


func (s *followServiceImpl) UnfollowUser(followerID, followedID uint) error {
	
	if followerID == followedID {
		return errors.New("cannot unfollow yourself")
	}

	
	if !s.followRepo.IsFollowing(followerID, followedID) {
		return errors.New("not following this user")
	}

	return s.followRepo.Delete(followerID, followedID)
}


func (s *followServiceImpl) GetFollowersCount(userID uint) (*response.FollowersCountResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	count := s.followRepo.CountFollowers(userID)

	return &response.FollowersCountResponse{
		UserID:         user.UserID,
		UserName:       user.UserName,
		FollowersCount: count,
	}, nil
}


func (s *followServiceImpl) GetFollowersList(userID uint, order string) (*response.FollowersListResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	followers, err := s.followRepo.GetFollowers(userID, order)
	if err != nil {
		return nil, err
	}

	
	followersDTO := make([]response.UserDTO, len(followers))
	for i, f := range followers {
		followersDTO[i] = response.UserDTO{
			UserID:   f.UserID,
			UserName: f.UserName,
		}
	}

	return &response.FollowersListResponse{
		UserID:    user.UserID,
		UserName:  user.UserName,
		Followers: followersDTO,
	}, nil
}


func (s *followServiceImpl) GetFollowedList(userID uint, order string) (*response.FollowedListResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	followed, err := s.followRepo.GetFollowed(userID, order)
	if err != nil {
		return nil, err
	}

	
	followedDTO := make([]response.UserDTO, len(followed))
	for i, f := range followed {
		followedDTO[i] = response.UserDTO{
			UserID:   f.UserID,
			UserName: f.UserName,
		}
	}

	return &response.FollowedListResponse{
		UserID:   user.UserID,
		UserName: user.UserName,
		Followed: followedDTO,
	}, nil
}