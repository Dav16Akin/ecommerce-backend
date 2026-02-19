package service

import (
	"errors"

	model "github.com/Dav16Akin/ecommerce-rest-backend/internal/models"
	"github.com/Dav16Akin/ecommerce-rest-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUser(id uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	if user.Email == "" {
		return errors.New("Email is required")
	}

	if user.Name == "" {
		return errors.New("Name is required")
	}

	if user.PhoneNumber == "" {
		return errors.New("Phone Number is required")
	}

	if user.Password == "" {
		return errors.New("Password is required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return s.repo.CreateUser(user)
}

func (s *userService) GetUser(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}