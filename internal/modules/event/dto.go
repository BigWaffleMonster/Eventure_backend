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

type CreateEventRequest struct {
	Title       string     `json:"title" binding:"required,min=3,max=100"`
	Description string     `json:"description" binding:"required"`
	MaxCapacity *int       `json:"max_capacity"`
	Location    *string    `json:"location"`
	StartDate   time.Time  `json:"start_date" binding:"required"`
	EndDate     *time.Time `json:"end_date"`
	CategoryID  uuid.UUID  `json:"category_id" binding:""`
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
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Capacity    int        `json:"capcity"`
	MaxCapacity *int       `json:"max_capacity"`
	Location    *string    `json:"location"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	DateCreated time.Time  `json:"date_created"`
	DateUpdated time.Time  `json:"date_updated"`

	Category CategoryResponse
	Owner    OwnerResponse
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
