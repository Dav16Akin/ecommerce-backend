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
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	UpdatePassword(id uint, oldPassword string, newPassword string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidPassword = errors.New("invalid password")

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

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	rows, err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrUserNotFound
	}

	return nil
}

func (s *userService) UpdatePassword(id uint, oldPassword string, newPassword string) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return ErrUserNotFound
	}

	// Compare old password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return ErrInvalidPassword
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return s.repo.UpdateUser(user)
}
