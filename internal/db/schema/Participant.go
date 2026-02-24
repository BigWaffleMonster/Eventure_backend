package schema

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID      uuid.UUID `gorm:"type:uuid;index"`
	EventID     uuid.UUID `gorm:"type:uuid;index"`
	DateCreated time.Time `gorm:"autoCreateTime"`

	User  User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Event Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"`
}
