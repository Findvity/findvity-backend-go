package event

import "github.com/BRO3886/findvity-backend/pkg/entities"

//Service for `event`
type Service interface {
	FindByID(id string) (*entities.Event, error)
	CreateEvent(event *entities.Event) (*entities.Event, error)
}

//NewService Creates a new repo
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

type service struct {
	repo Repository
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s *service) CreateEvent(event *entities.Event) (*entities.Event, error) {
	return s.repo.CreateEvent(event)
}

func (s *service) FindByID(id string) (*entities.Event, error) {
	return s.repo.FindByID(id)
}
