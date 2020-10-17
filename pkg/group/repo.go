package group

import (
	"github.com/BRO3886/findvity-backend/pkg"
	uuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"
)

//Repository for `group`
type Repository interface {
	FindByID(id string) (*Group, error)
	CreateGroup(group *Group) (*Group, error)
	AddMember(groupID, userID string) (*Group, error)
	UpdateGroup(group *Group) (*Group, error)
}

type repo struct {
	DB *gorm.DB
}

func (r *repo) FindByID(id string) (*Group, error) {
	group := &Group{}
	if err := r.DB.Where("id = ?", id).First(group).Error; err != nil {
		return nil, pkg.ErrNotFound
	}
	return group, nil
}

func (r *repo) CreateGroup(group *Group) (*Group, error) {
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

func (r *repo) AddMember(groupID, userID string) (*Group, error) {
	group := &Group{}
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

func (r *repo) UpdateGroup(group *Group) (*Group, error) {
	if err := r.DB.Model(&group).Updates(group).Error; err != nil {
		return nil, err
	}
	return group, nil
}
