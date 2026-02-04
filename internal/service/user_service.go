package service

import (
    "errors"
    
    "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
    "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository"
    "github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/utils"
)

type UserService interface {
    CreateUser(userName string, userType domain.UserType, password string) (*domain.User, error)
    GetUserByID(id uint) (*domain.User, error)
    GetAllUsers() ([]*domain.User, error)
    UpdateUser(id uint, userName string, userType domain.UserType, password string) (*domain.User, error)
    DeleteUser(id uint) error
    ValidatePassword(userID uint, password string) error
}

type userServiceImpl struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userServiceImpl{
        userRepo: userRepo,
    }
}

func (s *userServiceImpl) CreateUser(userName string, userType domain.UserType, password string) (*domain.User, error) {
    
    if userName == "" {
        return nil, errors.New("user name cannot be empty")
    }
    
    if userType != domain.UserTypeBuyer && userType != domain.UserTypeSeller && userType != domain.UserTypeBoth {
        return nil, errors.New("invalid user type")
    }
    
    // Validar senha
    if err := utils.IsPasswordValid(password); err != nil {
        return nil, err
    }
    
    // Hashear senha
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return nil, err
    }
    
    existingUser, _ := s.userRepo.FindByUsername(userName)
    if existingUser != nil {
        return nil, errors.New("username already exists")
    }
    
    user := &domain.User{
        UserName: userName,
        UserType: userType,
        Password: hashedPassword,
    }
    
    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *userServiceImpl) GetUserByID(id uint) (*domain.User, error) {
    return s.userRepo.FindByID(id)
}

func (s *userServiceImpl) GetAllUsers() ([]*domain.User, error) {
    return s.userRepo.FindAll()
}

func (s *userServiceImpl) UpdateUser(id uint, userName string, userType domain.UserType, password string) (*domain.User, error) {
    user, err := s.userRepo.FindByID(id)
    if err != nil {
        return nil, err
    }
    
    
    if userName != "" {
        user.UserName = userName
    }
    
    if userType != "" {
        if userType != domain.UserTypeBuyer && userType != domain.UserTypeSeller && userType != domain.UserTypeBoth {
            return nil, errors.New("invalid user type")
        }
        user.UserType = userType
    }
    
    // Atualizar senha se fornecida
    if password != "" {
        if err := utils.IsPasswordValid(password); err != nil {
            return nil, err
        }
        
        hashedPassword, err := utils.HashPassword(password)
        if err != nil {
            return nil, err
        }
        
        user.Password = hashedPassword
    }
    
    if err := s.userRepo.Update(user); err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *userServiceImpl) DeleteUser(id uint) error {
    if !s.userRepo.ExistsByID(id) {
        return errors.New("user not found")
    }
    
    return s.userRepo.Delete(id)
}

// ValidatePassword verifica se a senha fornecida corresponde ao hash armazenado
func (s *userServiceImpl) ValidatePassword(userID uint, password string) error {
    user, err := s.userRepo.FindByID(userID)
    if err != nil {
        return errors.New("user not found")
    }
    
    return utils.ComparePassword(user.Password, password)
}