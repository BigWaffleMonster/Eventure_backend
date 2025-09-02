package user

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/interfaces"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/utils/requests"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	interfaces.IBaseRepository[User]
	GetByEmail(email string) (*User, results.Result)
	CreateUserSession(userID uuid.UUID, requestInfo requests.RequestInfo) (*uuid.UUID, results.Result)
	DeleteUserSession(sessionID uuid.UUID) results.Result
	GetUserSession(sessionID uuid.UUID) (*UserSession, results.Result)
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

func (r *userRepository) GetUserSession(sessionID uuid.UUID) (*UserSession, results.Result) {
	var session UserSession
	result := r.DB.Where("id = ?", sessionID).First(&session)

	if result.Error != nil {
		return nil, results.NewInternalError(result.Error.Error())
	}

	return &session, results.NewResultOk()
}

func (r *userRepository) CreateUserSession(userID uuid.UUID, requestInfo requests.RequestInfo) (*uuid.UUID, results.Result) {
	data := UserSession{
			ID : uuid.New(),
			UserID: userID,
			IPAddress: requestInfo.IP,
			UserAgent: requestInfo.UserAgent,
			Fingerprint: requestInfo.Fingerprint,
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
			CreatedAt: time.Now(),
		}

	err := r.DB.Create(data).Error

	if err != nil {
		return nil, results.NewInternalError(err.Error())
	}

	return &data.ID, results.NewResultOk()
}

func (r *userRepository) DeleteUserSession(sessionID uuid.UUID) results.Result {
	err := r.DB.Delete(&UserSession{}, sessionID).Error

	if err != nil {
		return results.NewInternalError(err.Error())
	}

	return results.NewResultOk()
}