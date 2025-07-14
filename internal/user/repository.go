package user

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*User, results.Result)
	Create(user *User) results.Result
	GetByID(id uuid.UUID) (*User, results.Result)
	Update(user *User) results.Result
	Delete(id uuid.UUID) results.Result
	GetRefreshToken(refreshToken string) results.Result
	SetRefreshToken(userID uuid.UUID, token string) results.Result
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) GetByEmail(email string) (*User, results.Result) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &user, results.NewResultOk()
}

func (r *userRepository) Create(user *User) results.Result {
	err := r.DB.Create(user).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *userRepository) GetByID(id uuid.UUID) (*User, results.Result) {
	var user User
	result := r.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &user, results.NewResultOk()
}

func (r *userRepository) Update(user *User) results.Result {
	err := r.DB.Save(user).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *userRepository) Delete(id uuid.UUID) results.Result {
	err := r.DB.Delete(&User{}, id).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}

func (r *userRepository) GetRefreshToken(refreshToken string) results.Result {
	var token UserRefreshToken
	result := r.DB.Where("token = ?", refreshToken).First(&token)

	if result.Error != nil {
		return results.NewInternalError(result.Error.Error())
	}

	return results.NewResultOk()
}

func (r *userRepository) SetRefreshToken(userID uuid.UUID, token string) results.Result {
	var data UserRefreshToken

	result := r.DB.Where(UserRefreshToken{UserID: userID}).FirstOrCreate(&data, UserRefreshToken{
		UserID:        userID,
		RefsreshToken: token,
	})

	if result.RowsAffected == 0 {
		data.RefsreshToken = token
		err := r.DB.Save(data).Error

		if err != nil {
			return results.NewInternalError(err.Error())
		}

		return results.NewResultOk()
	} else {
		return results.NewInternalError(result.Error.Error())
	}
}
