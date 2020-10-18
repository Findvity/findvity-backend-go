package user

import (
	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `user`
type Repository interface {
	FindByID(id string) (*entities.User, error)

	FindByUsername(username string) (*entities.User, error)

	CreateUser(user *entities.User) (*entities.User, error)

	DoesUsernameExist(username string) (bool, error)
}

type repo struct {
	DB *gorm.DB
}

//NewRepo for user
func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

func (r *repo) FindByID(id string) (*entities.User, error) {
	user := &entities.User{}
	r.DB.Where("id = ?", id).First(user)
	if user.Name == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) FindByUsername(username string) (*entities.User, error) {
	user := &entities.User{}
	r.DB.Where("username = ?", username).First(user)
	if user.Name == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) CreateUser(user *entities.User) (*entities.User, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	user.ID = uid.String()
	result := r.DB.Create(user)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) DoesUsernameExist(username string) (bool, error) {
	user := &entities.User{}
	if err := r.DB.Where("username = ?", username).First(user).Error; err != nil {
		return false, pkg.ErrNotFound
	}
	return true, nil
}
