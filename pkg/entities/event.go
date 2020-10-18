package entities

import (
	"time"
)

//Event for group
type Event struct {
	Base
	Name      string    `json:"name"`
	GroupID   string    `json:"group_id"`
	PhotoURL  string    `json:"cover_url"`
	ExpiresAt time.Time `json:"expires_at"`
}
