package participant

import (
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/event"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/google/uuid"
)

type Participant struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid"`
	UserID      uuid.UUID `gorm:"type:uuid"`
	User        user.User   `gorm:"foreignKey:UserID"`
	EventID     uuid.UUID `gorm:"type:uuid"`
	Event       event.Event `gorm:"foreignKey:EventID"`
	Status      string    `sql:"type:ENUM('Yes', 'No', 'Maybe')"`
	HasAccess   bool
	Ticket      string
	DateCreated time.Time `gorm:"autoCreateTime"`
}
