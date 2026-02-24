package schema

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`

	Login    string `gorm:"index:unique" json:"login"`
	Email    string `gorm:"index:unique;not null" json:"email"`
	Password string `json:"password"`

	IsEmailConfirmed bool `gorm:"default:false" json:"is_email_confirmed"`

	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `gorm:"autoUpdateTime" json:"date_updated"`
}
