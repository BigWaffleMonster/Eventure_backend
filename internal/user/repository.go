package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetByID(id uuid.UUID) (*User, error) {
	var user User
	result := r.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Update(data *User) error {
	return r.DB.Save(data).Error
}

func (r *UserRepository) Remove(id uuid.UUID) error {
	return r.DB.Delete(&User{}, id).Error
}
