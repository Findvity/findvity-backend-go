package handlers

import (
	"net/http"

	"github.com/BRO3886/findvity-backend/api/middleware"
	"github.com/BRO3886/findvity-backend/pkg/entities"
	"github.com/BRO3886/findvity-backend/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func register(svc user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := &entities.User{}

		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg":   "unable to parse json",
				"error": err.Error(),
			})
		}

		doesExist, _ := svc.DoesUsernameExist(user.Username)

		if doesExist {
			return ctx.Status(http.StatusConflict).JSON(fiber.Map{
				"msg": "user with that username already exists",
			})
		}

		user, err := svc.Register(user)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "unable to create user",
				"error": err.Error(),
			})
		}

		token, err := middleware.CreateToken(user.Username, user.ID)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"msg":   "unable to create token for user",
				"error": err.Error(),
			})
		}

		user.Password = ""

		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"msg":   "user created",
			"token": token,
			"user":  *user,
		})
	}
}

func login(svc user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := &entities.User{}
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg":   "unable to parse json",
				"error": err.Error(),
			})
		}
		user, err := svc.Login(user.Username, user.Password)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "incorrect credentials",
				"error": err.Error(),
			})
		}
		token, err := middleware.CreateToken(user.Username, user.ID)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"msg":   "unable to create token for user",
				"error": err.Error(),
			})
		}

		user.Password = ""

		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"msg":   "login successful",
			"token": token,
			// "user":  *user,
		})
	}
}

//UserEndpoints manage user endpoints
func UserEndpoints(app *fiber.App, svc user.Service) {
	usrGroup := app.Group("/api/user")
	usrGroup.Post("/register", register(svc))
	usrGroup.Post("/login", login(svc))

}
