package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/request"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository/mocks"
)

func TestGetFollowedPosts_OrderDateAsc(t *testing.T) {
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)
	
	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)
	
	userID := uint(1)
	order := "date_asc"
	
	userRepoMock.On("ExistsByID", userID).Return(true)
	
	followedUsers := []*domain.User{
		{UserID: 2, UserName: "maria"},
	}
	followRepoMock.On("GetFollowed", userID, "").Return(followedUsers, nil)
	
	posts := []*domain.Post{
		{
			PostID: 1,
			UserID: 2,
			Date:   time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 1, ProductName: "Produto 1"},
		},
		{
			PostID: 2,
			UserID: 2,
			Date:   time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 2, ProductName: "Produto 2"},
		},
		{
			PostID: 3,
			UserID: 2,
			Date:   time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 3, ProductName: "Produto 3"},
		},
	}
	
	postRepoMock.On("GetPostsByUsersInDateRange", 
		mock.AnythingOfType("[]uint"), 
		mock.AnythingOfType("time.Time"), 
		mock.AnythingOfType("time.Time"), 
		order,
	).Return(posts, nil)
	
	result, err := service.GetFollowedPosts(userID, order)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Posts))
	
	assert.Equal(t, "10-01-2026", result.Posts[0].Date) // mais antigo
	assert.Equal(t, "15-01-2026", result.Posts[1].Date)
	assert.Equal(t, "20-01-2026", result.Posts[2].Date) // mais recente
	
	postRepoMock.AssertExpectations(t)
}

func TestGetFollowedPosts_OrderDateDesc(t *testing.T) {
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)
	
	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)
	
	userID := uint(1)
	order := "date_desc"
	
	userRepoMock.On("ExistsByID", userID).Return(true)
	
	followedUsers := []*domain.User{
		{UserID: 2, UserName: "maria"},
	}
	followRepoMock.On("GetFollowed", userID, "").Return(followedUsers, nil)
	
	posts := []*domain.Post{
		{
			PostID: 3,
			UserID: 2,
			Date:   time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 3, ProductName: "Produto 3"},
		},
		{
			PostID: 2,
			UserID: 2,
			Date:   time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 2, ProductName: "Produto 2"},
		},
		{
			PostID: 1,
			UserID: 2,
			Date:   time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
			Product: domain.Product{ProductID: 1, ProductName: "Produto 1"},
		},
	}
	
	postRepoMock.On("GetPostsByUsersInDateRange", 
		mock.AnythingOfType("[]uint"), 
		mock.AnythingOfType("time.Time"), 
		mock.AnythingOfType("time.Time"), 
		order,
	).Return(posts, nil)
	
	result, err := service.GetFollowedPosts(userID, order)
	
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 3, len(result.Posts))
	
	assert.Equal(t, "20-01-2026", result.Posts[0].Date) // mais recente
	assert.Equal(t, "15-01-2026", result.Posts[1].Date)
	assert.Equal(t, "10-01-2026", result.Posts[2].Date) // mais antigo
	
	postRepoMock.AssertExpectations(t)
}

func TestCreatePost_InvalidDateFormat(t *testing.T) {
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)
	
	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)
	
	req := request.CreatePostRequest{
		UserID: 1,
		Date:   "2026-01-15", // formato ERRADO (deve ser dd-MM-yyyy)
		Product: request.ProductRequest{
			ProductName: "Teste",
			Type:        "Teste",
		},
		Category: 1,
		Price:    100,
	}
	
	userRepoMock.On("ExistsByID", uint(1)).Return(true)
	
	err := service.CreatePost(req)
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid date format")
	
	productRepoMock.AssertNotCalled(t, "Create")
	postRepoMock.AssertNotCalled(t, "Create")
}

func TestCreatePromoPost_InvalidDiscount(t *testing.T) {
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)
	
	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)
	
	req := request.CreatePromoPostRequest{
		UserID: 1,
		Date:   "15-01-2026",
		Product: request.ProductRequest{
			ProductName: "Teste",
			Type:        "Teste",
		},
		Category: 1,
		Price:    100,
		HasPromo: true,
		Discount: 150,
	}
	
	userRepoMock.On("ExistsByID", uint(1)).Return(true)
	
	err := service.CreatePromoPost(req)
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "discount must be between 0 and 100")
}