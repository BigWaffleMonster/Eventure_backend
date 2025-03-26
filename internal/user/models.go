package user

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/notification"
	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserName         string    `gorm:"unique"`
	Email            string    `gorm:"unique;not null"`
	Password         string
	DateCreated      time.Time
	DateUpdated      time.Time `gorm:"autoUpdateTime"`
	IsEmailConfirmed bool      `gorm:"default:false"`

	// Events        []event.Event `gorm:"many2many:participants;"`
	Notifications []notification.Notification

	// Event Event `gorm:"foreignKey:OwnerID"`
}

type UserRegisterInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}

type UserLoginInput struct {
	Email    string  `json:"email"`
	Password *string `json:"password"`
}
