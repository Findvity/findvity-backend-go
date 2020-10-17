package utils

import (
	"github.com/BRO3886/findvity-backend/api/handlers"
	"github.com/BRO3886/findvity-backend/pkg/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//MakeUserHandlers creates user-hanlders
func MakeUserHandlers(app *fiber.App, db *gorm.DB) {
	userRepo := user.NewRepo(db)
	userSvc := user.NewService(userRepo)
	handlers.UserEndpoints(app, userSvc)
}
