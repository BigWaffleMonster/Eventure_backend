package views

import (
	"time"

	"github.com/google/uuid"
)

// @description event information
type EventInfo struct {
	Title                string `json:"title" example:"My best birth day"`
	Description          string `json:"description" example:"My best birth day"`
	Location             string `json:"location" example:"My best home"`
	Private              bool `json:"private" default:"false"`
	StartDate            time.Time `json:"startAt" format:"date-time"`
	EndDate              time.Time `json:"endAt" format:"date-time"`
	CategoryID           uuid.UUID `json:"category" format:"uuid"`
}
