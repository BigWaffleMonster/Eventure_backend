package schema

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Capacity    int       `gorm:"default:0" json:"capacity"`
	MaxCapacity int       `json:"max_capacity"`
	Location    Location  `gorm:"type:jsonb;column:location" json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `gorm:"autoUpdateTime" json:"date_updated"`

	Cover *string `json:"cover"`

	CategoryID uuid.UUID `gorm:"type:uuid"`
	OwnerID    uuid.UUID `gorm:"type:uuid"`

	Owner    User     `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE"`
	Category Category `gorm:"foreignKey:CategoryID;constraint:OnDelete:SET NULL"`
}

type Location struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	PlaceID int     `json:"place_id"`
	Address *string `json:"address"`
}
