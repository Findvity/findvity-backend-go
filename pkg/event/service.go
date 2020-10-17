package event

//Service for `event`
type Service interface {
	FindByID(id string) (*Event, error)
	CreateEvent(event *Event) (*Event, error)
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

func (s *service) CreateEvent(event *Event) (*Event, error) {
	return s.repo.CreateEvent(event)
}

func (s *service) FindByID(id string) (*Event, error) {
	return s.repo.FindByID(id)
}
