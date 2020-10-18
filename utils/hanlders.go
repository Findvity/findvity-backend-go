package utils

import (
	"github.com/BRO3886/findvity-backend/api/handlers"
	"github.com/BRO3886/findvity-backend/pkg/group"
	"github.com/BRO3886/findvity-backend/pkg/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//MakeHandlers creates user-hanlders
func MakeHandlers(app *fiber.App, db *gorm.DB) {
	//user handlers
	userRepo := user.NewRepo(db)
	userSvc := user.NewService(userRepo)
	handlers.UserEndpoints(app, userSvc)

	// group handlers
	groupRepo := group.NewRepo(db)
	groupSvc := group.NewService(groupRepo)
	handlers.GroupEndpoints(app, groupSvc)
}
