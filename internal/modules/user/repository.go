package user

import (
	"net/http"

	"github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/BigWaffleMonster/Eventure_backend/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserById(id uuid.UUID) (*schema.User, error) {
	var user_raw schema.User

	result := r.db.
		Where("event_id = ?", id).
		Find(&user_raw)

	if result.Error != nil {
		return nil, utils.NewAppErrorWithErr(
			http.StatusInternalServerError,
			"Failed to fetch participants",
			result.Error,
		)
	}

	return &user_raw, nil
}

func (r *UserRepository) GetUserByEmail(email string) {
}
