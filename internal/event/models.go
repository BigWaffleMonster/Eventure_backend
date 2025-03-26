package event

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/category"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/google/uuid"
)

type Event struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid"`
	OwnerID              uuid.UUID `grom:"type:uuid"`
	Title                string    `gorm:"unique; not null"`
	Description          string    `gorm:"not null"`
	NumberOfParticipants int
	Location             string
	Private              bool `gorm:"default:false"`
	StartDate            time.Time
	EndDate              time.Time
	DateCreated          time.Time
	DateUpdated          time.Time `gorm:"autoUpdateTime"`

	User  user.User   `gorm:"foreignKey:OwnerID;unique"`
	Users []user.User `gorm:"many2many:participants"`

	CategoryID uuid.UUID         `gorm:"unique;type:uuid"`
	Category   category.Category `gorm:"foreignKey:CategoryID"`
}
