package user

import (
	"github.com/BRO3886/findvity-backend/pkg"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `user`
type Repository interface {
	FindByID(id string) (*User, error)

	FindByUsername(username string) (*User, error)

	CreateUser(user *User) (*User, error)

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

func (r *repo) FindByID(id string) (*User, error) {
	user := &User{}
	r.DB.Where("id = ?", id).First(user)
	if user.Name == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) FindByUsername(username string) (*User, error) {
	user := &User{}
	r.DB.Where("username = ?", username).First(user)
	if user.Name == "" {
		return nil, pkg.ErrNotFound
	}
	return user, nil
}

func (r *repo) CreateUser(user *User) (*User, error) {
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
	user := &User{}
	if err := r.DB.Where("username = ?", username).First(user).Error; err != nil {
		return false, pkg.ErrNotFound
	}
	return true, nil
}
