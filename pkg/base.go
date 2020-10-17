package pkg

import "time"

//Base struct for gorm
type Base struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"`
}
