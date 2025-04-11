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
	GetRefreshToken(refreshToken string) error
	SetRefreshToken(userID uuid.UUID, token string) error
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

func (r *userRepository) GetRefreshToken(refreshToken string) error {
	var token UserRefreshToken
	result := r.DB.Where("token = ?", refreshToken).First(&token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) SetRefreshToken(userID uuid.UUID, token string) error {
	var data UserRefreshToken

	result := r.DB.Where(UserRefreshToken{UserID: userID}).FirstOrCreate(&data, UserRefreshToken{
		UserID:        userID,
		RefsreshToken: token,
	})

	if result.RowsAffected == 0 {
		data.RefsreshToken = token
		return r.DB.Save(&data).Error
	} else {
		return result.Error
	}
}
