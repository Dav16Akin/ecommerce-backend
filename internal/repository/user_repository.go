package repository

import (
	model "github.com/Dav16Akin/ecommerce-rest-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User)  error
}

type userRepository struct {
	db *gorm.DB
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}
