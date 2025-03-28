package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid"`
	OwnerID              uuid.UUID `grom:"type:uuid"`
	Title                string    `gorm:"not null"`
	Description          string    `gorm:"not null"`
	MaxQtyParticipants   int
	Location             string
	Private              bool `gorm:"default:false"`
	StartDate            time.Time
	EndDate              time.Time
	DateCreated          time.Time
	DateUpdated          time.Time `gorm:"autoUpdateTime"`

	//User  user.User   `gorm:"foreignKey:OwnerID;unique"`

	CategoryID uuid.UUID         `gorm:"type:uuid"`
	//Category   category.Category `gorm:"foreignKey:CategoryID"`
}

// @description event information
type EventInput struct {
	Title                string `json:"title" example:"My best birth day"`
	//OwnerID			     uuid.UUID
	Description          string `json:"description" example:"My best birth day"`
	MaxQtyParticipants   int    `json:"maxQtyParticipants" example:"30"`
	Location             string `json:"location" example:"My best home"`
	Private              bool   `json:"private" default:"false"`
	StartDate            time.Time `json:"startAt" format:"date-time"`
	EndDate              time.Time `json:"endAt" format:"date-time"`
	CategoryID           uuid.UUID `json:"category" format:"uuid"`
}

type EventView struct {
	ID                   uuid.UUID `json:"id"`
	OwnerID              uuid.UUID `json:"ownerId"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	MaxQtyParticipants   int       `json:"maxQtyParticipants"`
	Location             string    `json:"location"`
	Private              bool      `json:"private"`
	StartDate            time.Time `json:"startDate"`
	EndDate              time.Time `json:"endDate"`
	DateCreated          time.Time `json:"dateCreated"`
	DateUpdated          time.Time `json:"dateUpdated"`
	CategoryID           uuid.UUID `json:"categoryID"`
}