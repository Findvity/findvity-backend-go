package group

import "github.com/BRO3886/findvity-backend/pkg/entities"

//Service interface for Group
type Service interface {
	GetRepo() Repository
	GetGroupByID(id string) (*entities.Group, error)
	CreateGroup(group *entities.Group) (*entities.Group, error)
	AddMember(groupID, userID string) (*entities.Group, error)
	UpdateGroup(group *entities.Group) (*entities.Group, error)
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

func (s *service) GetGroupByID(id string) (*entities.Group, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateGroup(group *entities.Group) (*entities.Group, error) {
	return s.repo.CreateGroup(group)
}

func (s *service) AddMember(groupID, userID string) (*entities.Group, error) {
	return s.repo.AddMember(groupID, userID)
}

func (s *service) UpdateGroup(group *entities.Group) (*entities.Group, error) {
	return s.repo.UpdateGroup(group)
}
