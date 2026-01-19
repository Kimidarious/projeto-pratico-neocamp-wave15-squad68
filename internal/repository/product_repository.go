package repository

import (
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	FindByID(id uint) (*domain.Product, error)
}