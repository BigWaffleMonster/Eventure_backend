package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                   uuid.UUID `gorm:"primaryKey"`
	OwnerID              uuid.UUID
	Title                string `json:"email" gorm:"unique; not null"`
	Description          string `json:"password" gorm:"not noll"`
	NumberOfParticipants int
	Location             string
	Private              bool `gorm:"default:false"`
	StartDate            time.Time
	EndDate              time.Time
	DateCreated          time.Time
	DateUpdated          time.Time `gorm:"autoCreateTime"`

	User  User   `gorm:"foreignKey:OwnerID;unique"`
	Users []User `gorm:"many2many:participants;"`

	CategoryID uuid.UUID `gorm:"unique"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
}
