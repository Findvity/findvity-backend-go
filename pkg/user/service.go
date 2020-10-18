package user

import (
	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

//Service interface for user
type Service interface {
	GetRepo() Repository
	GetUserByID(id string) (*entities.User, error)
	DoesUsernameExist(username string) (bool, error)
	Register(user *entities.User) (*entities.User, error)
	Login(username, password string) (*entities.User, error)
}

type service struct {
	repo Repository
}

//NewService Creates a new repo
func NewService(r Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s *service) GetUserByID(id string) (*entities.User, error) {
	return s.repo.FindByID(id)
}

func (s *service) DoesUsernameExist(username string) (bool, error) {
	return s.repo.DoesUsernameExist(username)
}

func (s *service) Register(user *entities.User) (*entities.User, error) {
	validate, err := validate(user)
	if !validate {
		return nil, err
	}
	pass, err := hashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pass

	return s.repo.CreateUser(user)
}

func (s *service) Login(username, password string) (*entities.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(password, user.Password) {
		return user, nil
	}
	return nil, pkg.ErrNotFound
}

func validate(user *entities.User) (bool, error) {
	if len(user.Password) < 6 || len(user.Password) > 60 {
		return false, pkg.ErrPassword
	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
