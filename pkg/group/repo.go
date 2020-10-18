package group

import (
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
	group := &entities.Group{}
	if err := r.DB.Where("id = ?", groupID).First(group).Error; err != nil {
		return nil, pkg.ErrNotFound
	}
	group.MemberCount++
	group, err := r.UpdateGroup(group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (r *repo) UpdateGroup(group *entities.Group) (*entities.Group, error) {
	if err := r.DB.Model(&group).Updates(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
