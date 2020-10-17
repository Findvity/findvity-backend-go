package user

import (
	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/group"
)

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
	pkg.Base
	Name          string        `json:"name"`
	Username      string        `json:"username"`
	Phone         string        `json:"phone"`
	Age           int           `json:"age"`
	Sex           Gender        `json:"gen	der"`
	Password      string        `json:"password"`
	ProfileImgURL string        `json:"profile_img_url"`
	Verified      bool          `json:"verified"`
	Tags          string        `json:"tags"`
	Groups        []group.Group `json:"groups"`
}
