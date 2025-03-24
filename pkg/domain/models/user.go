package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserName         string    `json:"userName" gorm:"unique; not null"`
	Email            string    `json:"email" gorm:"unique; not null"`
	Password         string    `json:"password" gorm:"not noll"`
	DateCreated      time.Time
	DateUpdated      time.Time `gorm:"autoCreateTime"`
	IsEmailConfirmed bool      `gorm:"default:false"`

	Events        []Event `gorm:"many2many:participants;"`
	Notifications []Notification

	//Event Event `gorm:"foreignKey:OwnerID"`
}
