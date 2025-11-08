package user

import (
	"context"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	interfaces.IBaseRepository[User]
	GetByEmail(ctx context.Context, email string) (*User, results.Result)
	CreateUserSession(ctx context.Context, userID uuid.UUID) (*uuid.UUID, results.Result)
	DeleteUserSession(ctx context.Context, sessionID uuid.UUID) results.Result
	GetUserSession(ctx context.Context, sessionID uuid.UUID) (*UserSession, results.Result)
}

type userRepository struct {
	repository.BaseRepository[User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		repository.BaseRepository[User]{DB: db},
	}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*User, results.Result) {
	var user User
	result := r.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &user, results.NewResultOk()
}

func (r *userRepository) GetUserSession(ctx context.Context, sessionID uuid.UUID) (*UserSession, results.Result) {
	var session UserSession
	result := r.DB.Where("id = ?", sessionID).First(&session)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &session, results.NewResultOk()
}

func (r *userRepository) CreateUserSession(ctx context.Context, userID uuid.UUID) (*uuid.UUID, results.Result) {

	requestInfoIP, _ := helpers.GetIP(ctx)
	requestInfoUserAgent, _:= helpers.GetUserAgent(ctx)
	requestInfoFingerprint, _ := helpers.GetFingerprint(ctx)

	data := UserSession{
			ID : uuid.New(),
			UserID: userID,
			IPAddress: requestInfoIP,
			UserAgent: requestInfoUserAgent,
			Fingerprint: requestInfoFingerprint,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			CreatedAt: time.Now(),
		}

	err := r.DB.Create(&data).Error

	if err != nil {
		return nil, results.NewInternalError(err.Error())
	}

	return &data.ID, results.NewResultOk()
}

func (r *userRepository) DeleteUserSession(ctx context.Context, sessionID uuid.UUID) results.Result {
	err := r.DB.Delete(&UserSession{}, sessionID).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}