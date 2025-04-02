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

	Notifications []notification.Notification
}

type UserGetView struct {
	ID               uuid.UUID `json:"id"`
	UserName         string    `json:"userName"`
	Email            string    `json:"email"`
	DateCreated      time.Time `json:"dateCreated"`
	DateUpdated      time.Time `json:"dateUpdated"`
	IsEmailConfirmed bool      `json:"isEmailConfirmed"`
}

type UserUpdateInput struct {
	UserName         *string `json:"userName"`
	Email            *string `json:"email"`
	IsEmailConfirmed *bool   `json:"isEmailConfirmed"`
	Password         *string `json:"password"`
}

type UserUpdateView struct{}
