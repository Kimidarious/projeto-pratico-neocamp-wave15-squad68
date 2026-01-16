package repository

import (
    "errors"
    "gorm.io/gorm"
    "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
)

type userRepositoryImpl struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByID(id uint) (*domain.User, error) {
    var user domain.User
    result := r.db.First(&user, id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *userRepositoryImpl) FindAll() ([]*domain.User, error) {
    var users []*domain.User
    result := r.db.Find(&users)
    if result.Error != nil {
        return nil, result.Error
    }
    return users, nil
}

func (r *userRepositoryImpl) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *userRepositoryImpl) ExistsByID(id uint) bool {
    var count int64
    r.db.Model(&domain.User{}).Where("user_id = ?", id).Count(&count)
    return count > 0
}

func (r *userRepositoryImpl) FindByUsername(username string) (*domain.User, error) {
    var user domain.User
    result := r.db.Where("user_name = ?", username).First(&user)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("user not found")
        }
        return nil, result.Error
    }
    return &user, nil
}

func (r *userRepositoryImpl) Update(user *domain.User) error {
    return r.db.Save(user).Error
}

func (r *userRepositoryImpl) Delete(id uint) error {
    return r.db.Delete(&domain.User{}, id).Error
}