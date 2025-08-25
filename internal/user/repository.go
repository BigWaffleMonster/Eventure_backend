package user

import (
	"errors"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	interfaces.IBaseRepository[User]
	GetRefreshToken(userID uuid.UUID) (*UserRefreshToken, results.Result)
	SetRefreshToken(userID uuid.UUID, token string) results.Result
	GetByEmail(email string) (*User, results.Result)
}

type userRepository struct {
	repository.BaseRepository[User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		repository.BaseRepository[User]{DB: db},
	}
}

func (r *userRepository) GetByEmail(email string) (*User, results.Result) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &user, results.NewResultOk()
}

func (r *userRepository) GetRefreshToken(userID uuid.UUID) (*UserRefreshToken, results.Result) {
	var token UserRefreshToken
	result := r.DB.Where("user_id = ?", userID).First(&token)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, results.NewResultOk()
	}

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &token, results.NewResultOk()
}

func (r *userRepository) SetRefreshToken(userID uuid.UUID, token string) results.Result {
	data := UserRefreshToken{
		UserID:        userID,
		RefreshToken: token,
	}

	err := r.DB.Create(data).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}