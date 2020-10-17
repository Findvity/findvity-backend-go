package group

import (
	"github.com/BRO3886/findvity-backend/pkg"
)

//Group for users
type Group struct {
	pkg.Base
	Name        string   `json:"name"`
	Tags        []string `json:"tags"`
	UserID      string   `json:"created_by"`
	MemberCount int      `json:"member_count"`
	MemberIds   []string `json:"member_ids"`
}
