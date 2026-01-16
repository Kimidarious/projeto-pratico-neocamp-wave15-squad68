package repository

import (
    "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type UserRepository interface {
    FindByID(id uint) (*domain.User, error)
    FindAll() ([]*domain.User, error)
    Create(user *domain.User) error
    Update(user *domain.User) error
    Delete(id uint) error                 
    ExistsByID(id uint) bool
    FindByUsername(username string) (*domain.User, error)
}