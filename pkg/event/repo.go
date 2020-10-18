package event

import (
	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `event`
type Repository interface {
	FindByID(id string) (*entities.Event, error)
	CreateEvent(event *entities.Event) (*entities.Event, error)
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) FindByID(id string) (*entities.Event, error) {
	event := &entities.Event{}
	if err := r.DB.Where("id = ?", id).First(event).Error; err != nil {
		return nil, pkg.ErrNotFound
	}
	return event, nil
}

func (r *repo) CreateEvent(event *entities.Event) (*entities.Event, error) {
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
