package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserName         string    `gorm:"index:unique"`
	Email            string    `gorm:"index:unique;not null"`
	Password         string
	DateCreated      time.Time
	DateUpdated      time.Time `gorm:"autoUpdateTime"`
	IsEmailConfirmed bool      `gorm:"default:false"`
}

type UserSession struct {
    ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
    UserID        uuid.UUID `gorm:"type:uuid;not null;index"`
    UserAgent     string    `gorm:"type:varchar(255);not null"`
    IPAddress     string    `gorm:"type:varchar(45);not null"` // IPv6 может быть до 45 символов
    Fingerprint   string    `gorm:"type:varchar(255);not null"`
    ExpiresAt     time.Time `gorm:"not null"`
    CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    
    // Опционально: связь с пользователем
    User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type UserView struct {
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
