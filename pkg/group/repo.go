package group

import (
	"log"

	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `group`
type Repository interface {
	FindByID(id string) (*entities.Group, error)
	CreateGroup(group *entities.Group) (*entities.Group, error)
	AddMember(groupID, userID string) (*entities.Group, error)
	UpdateGroup(group *entities.Group) (*entities.Group, error)
}

//NewRepo creates new repo
func NewRepo(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) FindByID(id string) (*entities.Group, error) {
	group := &entities.Group{}
	if err := r.DB.Where("id = ?", id).First(group).Error; err != nil {
		return nil, pkg.ErrNotFound
	}
	return group, nil
}

func (r *repo) CreateGroup(group *entities.Group) (*entities.Group, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	group.ID = uid.String()
	result := r.DB.Create(group)
	if result.Error != nil {
		return nil, pkg.ErrDatabase
	}
	return group, nil
}

func (r *repo) AddMember(groupID, userID string) (*entities.Group, error) {
	tx := r.DB.Begin()
	group := &entities.Group{}
	user := &entities.User{}

	//check if group exists
	if err := tx.Where("id=?", groupID).First(&group).Error; err != nil {
		tx.Rollback()
		log.Println("user not found")
		return nil, pkg.ErrDatabase
	}
	//check user
	if err := tx.Where("id=?", userID).First(&user).Error; err != nil {
		tx.Rollback()
		log.Println("user not found")
		return nil, pkg.ErrDatabase
	}
	if err := tx.Find(user).Association("Groups").Append(group).Error; err != nil {
		tx.Rollback()
		print(err)
		log.Println("unable to add to m2m")
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return group, nil
}

func (r *repo) UpdateGroup(group *entities.Group) (*entities.Group, error) {
	if err := r.DB.Model(&group).Updates(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
