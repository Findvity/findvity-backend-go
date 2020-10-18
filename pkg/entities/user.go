package entities

//Gender for user
type Gender string

const (
	//Male enum
	Male Gender = "Male"
	//Female enum
	Female = "Female"
	//NonBinary enum
	NonBinary = "Non-Binary"
	//Undisclosed enum
	Undisclosed = "Prefer not to disclose"
)

//User struct for user details
type User struct {
	Base
	Name          string   `json:"name"`
	Username      string   `json:"username"`
	Phone         string   `json:"phone"`
	Age           int      `json:"age"`
	Sex           Gender   `json:"gender"`
	Password      string   `json:"password"`
	ProfileImgURL string   `json:"profile_img_url"`
	Verified      bool     `json:"verified"`
	Tags          string   `json:"tags"`
	Groups        []*Group `json:"-" gorm:"many2many:user_groups;"`
}
