package handlers

import (
	"net/http"

	"github.com/BRO3886/findvity-backend/api"

	"github.com/BRO3886/findvity-backend/api/middleware"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	"github.com/BRO3886/findvity-backend/pkg/group"
	"github.com/gofiber/fiber/v2"
)

func create(svc group.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		group := &entities.Group{}
		// unmarshal body
		if err := ctx.BodyParser(&group); err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg":   "unable to parse json",
				"error": err.Error(),
			})
		}
		//get user id from header
		uid := middleware.GetUIDFromToken(ctx)
		if uid == "" {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"msg":   "Unable to parse token",
				"error": api.ErrInvalidToken,
			})
		}

		//create group
		group.UserID = uid
		group, err := svc.CreateGroup(group)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "unable to create group",
				"error": err.Error(),
			})
		}

		group, err = svc.AddMember(group.ID, uid)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "unable to add member to group",
				"error": err.Error(),
			})
		}

		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"msg":   "group",
			"group": *group,
		})
	}
}

//GroupEndpoints manages group endpoints
func GroupEndpoints(app *fiber.App, svc group.Service) {
	grpGroup := app.Group("/api/group")
	grpGroup.Use(middleware.BasicJWTAuth())
	{
		grpGroup.Post("/create", create(svc))
	}
	// usrGroup.Post("/api/user/login", login(svc))

}
