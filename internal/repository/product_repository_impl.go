package repository

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{db: db}
}

func (r *productRepositoryImpl) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepositoryImpl) FindByID(id uint) (*domain.Product, error) {
	var product domain.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return &product, nil
}