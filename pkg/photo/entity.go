package photo

import (
	"github.com/BRO3886/findvity-backend/pkg"
	"github.com/BRO3886/findvity-backend/pkg/group"
	"github.com/BRO3886/findvity-backend/pkg/user"
)

//Photo struct for feed photos
type Photo struct {
	pkg.Base
	PhotoURL string      `json:"img_url"`
	Likes    int         `json:"likes"`
	Group    group.Group `json:"group"`
	User     user.User   `json:"user"`
}
