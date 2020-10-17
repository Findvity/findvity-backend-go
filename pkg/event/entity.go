package event

import (
	"time"

	"github.com/BRO3886/findvity-backend/pkg"
)

//Event for group
type Event struct {
	pkg.Base
	Name      string    `json:"name"`
	GroupID   string    `json:"group_id"`
	PhotoURL  string    `json:"cover_url"`
	ExpiresAt time.Time `json:"expires_at"`
}
