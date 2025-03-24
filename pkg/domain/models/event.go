package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid"`
	OwnerID              uuid.UUID `gorm:"type:uuid"`
	Title                string `json:"email" gorm:"unique; not null"`
	Description          string `json:"password" gorm:"not noll"`
	//TODO: Лучше считать дигамически иначе нужно делать блокировку БД, причем желательно пессимистичную, 
	// тк оптимистичная будет в случае коллизий выдавать исклбчение
	//NumberOfParticipants int
	Location             string
	Private              bool `gorm:"default:false"`
	StartDate            time.Time
	EndDate              time.Time
	DateCreated          time.Time `gorm:"autoCreateTime"`
	DateUpdated          time.Time `gorm:"autoCreateTime"`

	User  User   `gorm:"foreignKey:OwnerID;type:uuid"`
	Users []User `gorm:"many2many:participants;"`

	CategoryID uuid.UUID `gorm:"type:uuid"`
	Category   Category  `gorm:"foreignKey:CategoryID"`
}
