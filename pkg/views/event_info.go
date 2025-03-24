package views

import (
	"time"

	"github.com/google/uuid"
)

type EventInfo struct {
	Title                string `json:"title"`
	Description          string `json:"description"`
	Location             string `json:"location"`
	Private              bool `json:"private"`
	StartDate            time.Time `json:"startAt"`
	EndDate              time.Time `json:"endAt"`
	CategoryID           uuid.UUID `json:"category"`
}
