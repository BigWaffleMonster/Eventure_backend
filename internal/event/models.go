package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid"`
	OwnerID              uuid.UUID `gorm:"type:uuid"`
	Title                string    `gorm:"not null"`
	Description          string    `gorm:"not null"`
	MaxQtyParticipants   int
	Location             string
	Private              bool      `gorm:"default:false"`
	StartDate            time.Time
	EndDate              time.Time
	DateCreated          time.Time
	DateUpdated          time.Time `gorm:"autoUpdateTime"`
	CategoryID           uuid.UUID `gorm:"type:uuid"`
}

// @description Событие
type EventInput struct {
	// Название
	Title                *string `json:"title" example:"My best birth day"`
	// Описание
	Description          *string `json:"description" example:"My best birth day"`
	// Максимальное кол-во участников
	MaxQtyParticipants   *int    `json:"maxQtyParticipants" example:"30"`
	// Локация
	Location             *string `json:"location" example:"My best home"`
	// Приватность
	Private              *bool   `json:"private" default:"false"`
	// Дата начала
	StartDate            *time.Time `json:"startAt" format:"date-time"`
	// Дата конца
	EndDate              *time.Time `json:"endAt" format:"date-time"`
	// Категория
	CategoryID           *uuid.UUID `json:"category" format:"uuid"`
}

// @description event view model
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