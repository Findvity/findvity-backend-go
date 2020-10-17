package event

import (
	"github.com/BRO3886/findvity-backend/pkg"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `event`
type Repository interface {
	FindByID(id string) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) FindByID(id string) (*Event, error) {
	event := &Event{}
	if err := r.DB.Where("id = ?", id).First(event).Error; err != nil {
		return nil, pkg.ErrNotFound
	}
	return event, nil
}

func (r *repo) CreateEvent(event *Event) (*Event, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	event.ID = uid.String()
	result := r.DB.Create(event)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return event, nil
}
