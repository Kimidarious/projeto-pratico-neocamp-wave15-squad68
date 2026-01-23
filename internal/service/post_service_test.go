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
	t.Log("US0006/US0009: Posts seguidos ordenados por data asc")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(1)
	order := "date_asc"

	userRepoMock.On("ExistsByID", userID).Return(true)

	followedUsers := []*domain.User{
		{UserID: 2, UserName: "maria", UserType: domain.UserTypeSeller},
	}
	followRepoMock.On("GetFollowed", userID, "").Return(followedUsers, nil)

	posts := []*domain.Post{
		{
			PostID:   1,
			UserID:   2,
			Date:     time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 1, ProductName: "Produto 1"}, // ← PONTEIRO
			Price:    100,
			Category: 1,
		},
		{
			PostID:   2,
			UserID:   2,
			Date:     time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 2, ProductName: "Produto 2"}, // ← PONTEIRO
			Price:    200,
			Category: 1,
		},
		{
			PostID:   3,
			UserID:   2,
			Date:     time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 3, ProductName: "Produto 3"}, // ← PONTEIRO
			Price:    300,
			Category: 1,
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
	assert.Equal(t, "10-01-2026", result.Posts[0].Date)
	assert.Equal(t, "15-01-2026", result.Posts[1].Date)
	assert.Equal(t, "20-01-2026", result.Posts[2].Date)

	postRepoMock.AssertExpectations(t)
}

func TestGetFollowedPosts_OrderDateDesc(t *testing.T) {
	t.Log("US0006/US0009: Posts seguidos ordenados por data desc")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(1)
	order := "date_desc"

	userRepoMock.On("ExistsByID", userID).Return(true)

	followedUsers := []*domain.User{
		{UserID: 2, UserName: "maria", UserType: domain.UserTypeSeller},
	}
	followRepoMock.On("GetFollowed", userID, "").Return(followedUsers, nil)

	posts := []*domain.Post{
		{
			PostID:   3,
			UserID:   2,
			Date:     time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 3, ProductName: "Produto 3"},
			Price:    300,
			Category: 1,
		},
		{
			PostID:   2,
			UserID:   2,
			Date:     time.Date(2026, 1, 15, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 2, ProductName: "Produto 2"},
			Price:    200,
			Category: 1,
		},
		{
			PostID:   1,
			UserID:   2,
			Date:     time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC),
			Product:  &domain.Product{ProductID: 1, ProductName: "Produto 1"},
			Price:    100,
			Category: 1,
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
	assert.Equal(t, "20-01-2026", result.Posts[0].Date)
	assert.Equal(t, "15-01-2026", result.Posts[1].Date)
	assert.Equal(t, "10-01-2026", result.Posts[2].Date)

	postRepoMock.AssertExpectations(t)
}

func TestCreatePost_InvalidDateFormat(t *testing.T) {
	t.Log("US0005: Criar publicação (data inválida)")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	req := request.CreatePostRequest{
		UserID: 1,
		Date:   "2026-01-15",
		Product: request.ProductRequest{
			ProductName: "Teste",
			Type:        "Teste",
		},
		Category: 1,
		Price:    100,
	}

	// Mock: usuário existe
	userRepoMock.On("ExistsByID", uint(1)).Return(true)

	// NÃO adicionar mock de Create porque deve falhar antes

	err := service.CreatePost(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid date format")

	productRepoMock.AssertNotCalled(t, "Create")
	postRepoMock.AssertNotCalled(t, "Create")
}

func TestCreatePromoPost_InvalidDiscount(t *testing.T) {
	t.Log("US0010: Criar publicação promocional (desconto inválido)")
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

func TestCreatePost_Success(t *testing.T) {
	t.Log("US0005: Criar publicação (sucesso)")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	req := request.CreatePostRequest{
		UserID: 1,
		Date:   "15-01-2026",
		Product: request.ProductRequest{
			ProductName: "Produto",
			Type:        "Tipo",
		},
		Category: 1,
		Price:    100,
	}

	userRepoMock.On("ExistsByID", uint(1)).Return(true)
	productRepoMock.On("Create", mock.AnythingOfType("*domain.Product")).Run(func(args mock.Arguments) {
		product := args.Get(0).(*domain.Product)
		product.ProductID = 10
	}).Return(nil)
	postRepoMock.On("Create", mock.AnythingOfType("*domain.Post")).Return(nil)

	err := service.CreatePost(req)

	assert.NoError(t, err)
	productRepoMock.AssertExpectations(t)
	postRepoMock.AssertExpectations(t)
}

func TestCreatePromoPost_Success(t *testing.T) {
	t.Log("US0010: Criar publicação promocional (sucesso)")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	req := request.CreatePromoPostRequest{
		UserID: 1,
		Date:   "15-01-2026",
		Product: request.ProductRequest{
			ProductName: "Promo",
			Type:        "Tipo",
		},
		Category: 1,
		Price:    100,
		HasPromo: true,
		Discount: 10,
	}

	userRepoMock.On("ExistsByID", uint(1)).Return(true)
	productRepoMock.On("Create", mock.AnythingOfType("*domain.Product")).Run(func(args mock.Arguments) {
		product := args.Get(0).(*domain.Product)
		product.ProductID = 11
	}).Return(nil)
	postRepoMock.On("Create", mock.AnythingOfType("*domain.Post")).Return(nil)

	err := service.CreatePromoPost(req)

	assert.NoError(t, err)
	productRepoMock.AssertExpectations(t)
	postRepoMock.AssertExpectations(t)
}

func TestGetFollowedPosts_EmptyFollowed(t *testing.T) {
	t.Log("US0006: Posts seguidos (sem vendedores seguidos)")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(1)
	order := "date_desc"

	followRepoMock.On("GetFollowed", userID, "").Return([]*domain.User{}, nil)

	result, err := service.GetFollowedPosts(userID, order)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result.Posts))
}

func TestCountPromoProducts_Success(t *testing.T) {
	t.Log("US0011: Contagem de produtos promocionais")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(1)
	user := &domain.User{
		UserID:   userID,
		UserName: "maria",
		UserType: domain.UserTypeSeller,
	}

	userRepoMock.On("FindByID", userID).Return(user, nil)
	postRepoMock.On("CountPromoPostsByUser", userID).Return(int64(3))

	result, err := service.CountPromoProducts(userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, int64(3), result.PromoProductsCount)
}

func TestGetPromoPostsByUser_Success(t *testing.T) {
	t.Log("US0012: Lista de produtos promocionais do vendedor")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(1)

	userRepoMock.On("ExistsByID", userID).Return(true)

	posts := []*domain.Post{
		{
			PostID:   1,
			UserID:   userID,
			Date:     time.Date(2026, 1, 20, 0, 0, 0, 0, time.UTC),
			HasPromo: true,
			Discount: 20,
			Product:  &domain.Product{ProductID: 1, ProductName: "Produto 1"},
		},
	}

	postRepoMock.On("GetPromoPostsByUser", userID).Return(posts, nil)

	result, err := service.GetPromoPostsByUser(userID)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Posts))
	assert.Equal(t, "20-01-2026", result.Posts[0].Date)
	assert.Equal(t, true, result.Posts[0].HasPromo)
	assert.Equal(t, float64(20), result.Posts[0].Discount)
}

func TestGetPromoPostsByUser_UserNotFound(t *testing.T) {
	t.Log("US0012: Lista de promo (usuário não encontrado)")
	postRepoMock := new(mocks.PostRepositoryMock)
	productRepoMock := new(mocks.ProductRepositoryMock)
	userRepoMock := new(mocks.UserRepositoryMock)
	followRepoMock := new(mocks.FollowRepositoryMock)

	service := NewPostService(postRepoMock, productRepoMock, userRepoMock, followRepoMock)

	userID := uint(999)

	userRepoMock.On("ExistsByID", userID).Return(false)

	result, err := service.GetPromoPostsByUser(userID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "user not found")
}
