package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"primaryKey"`
	UserName         string    `json:"userName" gorm:"unique"`
	Email            string    `json:"email" gorm:"unique; not null"`
	Password         string    `json:"password"`
	DateCreated      time.Time
	DateUpdated      time.Time `gorm:"autoCreateTime"`
	IsEmailConfirmed bool      `gorm:"default:false"`

	Events        []Event `gorm:"many2many:participants;"`
	Notifications []Notification

	// Event Event `gorm:"foreignKey:OwnerID"`
}
