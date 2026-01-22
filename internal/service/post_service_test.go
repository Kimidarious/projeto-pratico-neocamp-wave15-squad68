package service

import (
	"errors"
	"time"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/request"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/dto/response"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository"
)

type PostService interface {
	CreatePost(req request.CreatePostRequest) error
	CreatePromoPost(req request.CreatePromoPostRequest) error
	GetFollowedPosts(userID uint, order string) (*response.PostListResponse, error)
	CountPromoProducts(userID uint) (*response.PromoCountResponse, error)
	GetPromoPostsByUser(userID uint) (*response.PostListResponse, error)
}

type postServiceImpl struct {
	postRepo    repository.PostRepository
	productRepo repository.ProductRepository
	userRepo    repository.UserRepository
	followRepo  repository.FollowRepository
}

func NewPostService(
	postRepo repository.PostRepository,
	productRepo repository.ProductRepository,
	userRepo repository.UserRepository,
	followRepo repository.FollowRepository,
) PostService {
	return &postServiceImpl{
		postRepo:    postRepo,
		productRepo: productRepo,
		userRepo:    userRepo,
		followRepo:  followRepo,
	}
}


func (s *postServiceImpl) CreatePost(req request.CreatePostRequest) error {
	
	if !s.userRepo.ExistsByID(req.UserID) {
		return errors.New("user not found")
	}

	
	product := &domain.Product{
		ProductName: req.Product.ProductName,
		Type:        req.Product.Type,
		Brand:       req.Product.Brand,
		Color:       req.Product.Color,
		Notes:       req.Product.Notes,
	}

	if err := s.productRepo.Create(product); err != nil {
		return err
	}

	
	date, err := time.Parse("02-01-2006", req.Date)
	if err != nil {
		return errors.New("invalid date format, expected dd-MM-yyyy")
	}

	
	post := &domain.Post{
		UserID:    req.UserID,
		ProductID: product.ProductID,
		Date:      date,
		Category:  req.Category,
		Price:     req.Price,
		HasPromo:  false,
		Discount:  0,
	}

	return s.postRepo.Create(post)
}


func (s *postServiceImpl) CreatePromoPost(req request.CreatePromoPostRequest) error {
	
	if !s.userRepo.ExistsByID(req.UserID) {
		return errors.New("user not found")
	}

	
	if req.Discount < 0 || req.Discount > 100 {
		return errors.New("discount must be between 0 and 100")
	}

	
	product := &domain.Product{
		ProductName: req.Product.ProductName,
		Type:        req.Product.Type,
		Brand:       req.Product.Brand,
		Color:       req.Product.Color,
		Notes:       req.Product.Notes,
	}

	if err := s.productRepo.Create(product); err != nil {
		return err
	}

	
	date, err := time.Parse("02-01-2006", req.Date)
	if err != nil {
		return errors.New("invalid date format, expected dd-MM-yyyy")
	}

	
	post := &domain.Post{
		UserID:    req.UserID,
		ProductID: product.ProductID,
		Date:      date,
		Category:  req.Category,
		Price:     req.Price,
		HasPromo:  req.HasPromo,
		Discount:  req.Discount,
	}

	return s.postRepo.Create(post)
}


func (s *postServiceImpl) GetFollowedPosts(userID uint, order string) (*response.PostListResponse, error) {
	
	followed, err := s.followRepo.GetFollowed(userID, "")
	if err != nil {
		return nil, err
	}

	if len(followed) == 0 {
		return &response.PostListResponse{
			UserID: userID,
			Posts:  []response.PostDTO{},
		}, nil
	}

	
	followedIDs := make([]uint, len(followed))
	for i, u := range followed {
		followedIDs[i] = u.UserID
	}

	
	endDate := time.Now()
	startDate := endDate.AddDate(0, 0, -14)

	
	posts, err := s.postRepo.GetPostsByUsersInDateRange(followedIDs, startDate, endDate, order)
	if err != nil {
		return nil, err
	}

	
	postsDTO := make([]response.PostDTO, len(posts))
	for i, p := range posts {
		postsDTO[i] = response.PostDTO{
			PostID:   p.PostID,
			UserID:   p.UserID,
			Date:     p.Date.Format("02-01-2006"),
			Category: p.Category,
			Price:    p.Price,
			HasPromo: p.HasPromo,
			Discount: p.Discount,
			Product: response.ProductDTO{
				ProductID:   p.Product.ProductID,
				ProductName: p.Product.ProductName,
				Type:        p.Product.Type,
				Brand:       p.Product.Brand,
				Color:       p.Product.Color,
				Notes:       p.Product.Notes,
			},
		}
	}

	return &response.PostListResponse{
		UserID: userID,
		Posts:  postsDTO,
	}, nil
}


func (s *postServiceImpl) CountPromoProducts(userID uint) (*response.PromoCountResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	count := s.postRepo.CountPromoPostsByUser(userID)

	return &response.PromoCountResponse{
		UserID:             user.UserID,
		UserName:           user.UserName,
		PromoProductsCount: count,
	}, nil
}

func (s *postServiceImpl) GetPromoPostsByUser(userID uint) (*response.PostListResponse, error) {
	
	if !s.userRepo.ExistsByID(userID) {
		return nil, errors.New("user not found")
	}

	posts, err := s.postRepo.GetPromoPostsByUser(userID)
	if err != nil {
		return nil, err
	}

	postsDTO := make([]response.PostDTO, len(posts))
	for i, p := range posts {
		postsDTO[i] = response.PostDTO{
			PostID:   p.PostID,
			UserID:   p.UserID,
			Date:     p.Date.Format("02-01-2006"),
			Category: p.Category,
			Price:    p.Price,
			HasPromo: p.HasPromo,
			Discount: p.Discount,
			Product: response.ProductDTO{
				ProductID:   p.Product.ProductID,
				ProductName: p.Product.ProductName,
				Type:        p.Product.Type,
				Brand:       p.Product.Brand,
				Color:       p.Product.Color,
				Notes:       p.Product.Notes,
			},
		}
	}

	return &response.PostListResponse{
		UserID: userID,
		Posts:  postsDTO,
	}, nil
}