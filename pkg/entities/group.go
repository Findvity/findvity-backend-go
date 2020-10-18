package entities

//Group for users
type Group struct {
	Base
	Name    string `json:"name"`
	Tags    string `json:"tags"`
	UserID  string `json:"created_by"`
	Members []User `json:"members" gorm:"many2many:user_groups;"`
}
