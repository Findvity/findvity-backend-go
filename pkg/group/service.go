package group

//Service interface for Group
type Service interface {
	GetRepo() Repository
	GetGroupByID(id string) (*Group, error)
	CreateGroup(group *Group) (*Group, error)
	AddMember(groupID, userID string) (*Group, error)
	UpdateGroup(group *Group) (*Group, error)
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

func (s *service) GetGroupByID(id string) (*Group, error) {
	return s.repo.FindByID(id)
}

func (s *service) CreateGroup(group *Group) (*Group, error) {
	return s.repo.CreateGroup(group)
}

func (s *service) AddMember(groupID, userID string) (*Group, error) {
	return s.repo.AddMember(groupID, userID)
}

func (s *service) UpdateGroup(group *Group) (*Group, error) {
	return s.repo.UpdateGroup(group)
}
