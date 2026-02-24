package auth

import (
	"errors"

	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByID(id uuid.UUID) (*schema.User, error) {
	var user schema.User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) GetUserByLogin(login string) (*schema.User, error) {
	var existingUser schema.User

	if err := r.db.Where("email = ? OR login = ?", login, login).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &existingUser, nil
}

func (r *AuthRepository) CreateUser(newUser *schema.User) error {
	if err := r.db.Create(&newUser).Error; err != nil {
		// Проверка на уникальность email (на уровне БД)
		return err
		// ErrorResponse{
		// 	Error:   "create_failed",
		// 	Message: "Не удалось создать пользователя",
		// }
	}

	return nil
}
