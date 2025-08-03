package service

import (
	"errors"
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"go-backend/internal/utils"

	"github.com/google/uuid"
)

type UserService interface {
	Register(req *models.UserRegisterRequest) (*models.UserResponse, error)
	Login(req *models.UserLoginRequest) (*models.LoginResponse, error)
	GetCurrent(username string) (*models.UserResponse, error)
	Update(username string, req *models.UserUpdateRequest) (*models.UserResponse, error)
	Logout(username string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(req *models.UserRegisterRequest) (*models.UserResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	// Check if username already exists
	count, err := s.userRepo.CountByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("Username already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (s *userService) Login(req *models.UserLoginRequest) (*models.LoginResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Username or password wrong")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("Username or password wrong")
	}

	// Generate token
	token := uuid.New().String()
	user.Token = &token

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
	}, nil
}

func (s *userService) GetCurrent(username string) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not found")
	}

	return &models.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (s *userService) Update(username string, req *models.UserUpdateRequest) (*models.UserResponse, error) {
	if err := utils.ValidateStruct(req); err != nil {
		return nil, err
	}

	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user is not found")
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashPassword(*req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		Username: user.Username,
		Name:     user.Name,
	}, nil
}

func (s *userService) Logout(username string) error {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user is not found")
	}

	user.Token = nil
	return s.userRepo.Update(user)
}