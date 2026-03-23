package event

import (
	"time"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type OwnerResponse struct {
	ID    uuid.UUID `json:"id"`
	Login string    `json:"login"`
	Email string    `json:"email"`
}

type CreateEventForm struct {
	Title       string    `form:"title" binding:"required,min=3,max=100"`
	Description string    `form:"description" binding:"required"`
	CategoryID  string    `form:"category" binding:"required"`
	StartDate   time.Time `form:"startDate" time_format:"2006-01-02T15:04:05Z07:00" binding:"required"`
	EndDate     time.Time `form:"endDate" time_format:"2006-01-02T15:04:05Z07:00" binding:"required"`
	MaxCapacity int       `form:"maxCapacity"`

	LocationLat     float64 `form:"location[lat]" binding:"required"`
	LocationLng     float64 `form:"location[lng]" binding:"required"`
	LocationPlaceID int     `form:"location[place_id]" binding:"required"`
	LocationAddress string  `form:"location[address]"`
}

type CreateEventRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required"`

	MaxCapacity int `json:"maxCapacity"`

	Location location `json:"location" binding:"required"`

	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`

	CategoryID uuid.UUID `json:"category" binding:"required"`
}

type UpdateEventRequest struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Capacity    *int       `json:"capacity"`
	MaxCapacity *int       `json:"max_capacity"`
	Location    *string    `json:"location"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	CategoryID  *uuid.UUID `json:"category_id"`
}

type EventResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Capacity    int       `json:"capcity"`
	MaxCapacity int       `json:"max_capacity"`
	Location    location  `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`

	Cover *string `json:"cover"`

	Category CategoryResponse
	Owner    OwnerResponse
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type location struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	PlaceID int     `json:"place_id"`
	Address *string `json:"address"`
}
