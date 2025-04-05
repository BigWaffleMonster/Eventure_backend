package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*User, error)
	Create(user *User) error
	GetByID(id uuid.UUID) (*User, error)
	Update(data *User) error
	Remove(id uuid.UUID) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *userRepository) GetByID(id uuid.UUID) (*User, error) {
	var user User
	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(data *User) error {
	return r.DB.Save(data).Error
}

func (r *userRepository) Remove(id uuid.UUID) error {
	return r.DB.Delete(&User{}, id).Error
}
