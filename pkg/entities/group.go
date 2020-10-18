package entities

//Group for users
type Group struct {
	Base
	Name   string  `json:"name"`
	Tags   string  `json:"tags"`
	UserID string  `json:"created_by"`
	Users  []*User `json:"-" gorm:"many2many:user_groups;"`
}
