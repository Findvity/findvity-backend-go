package handlers

import (
	"net/http"

	"github.com/BRO3886/findvity-backend/api/middleware"
	"github.com/BRO3886/findvity-backend/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func register(svc user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := &user.User{}

		if err := ctx.BodyParser(&user); err != nil {
			ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg":   "unable to parse json",
				"error": err.Error(),
			})
			return err
		}

		doesExist, _ := svc.DoesUsernameExist(user.Username)

		if doesExist {
			ctx.Status(http.StatusConflict).JSON(fiber.Map{
				"msg": "user with that username already exists",
			})
			return nil
		}

		user, err := svc.Register(user)

		if err != nil {
			ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "unable to create user",
				"error": err.Error(),
			})
			return err
		}

		token, err := middleware.CreateToken(user.Username, user.ID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"msg":   "unable to create token for user",
				"error": err.Error(),
			})
			return err
		}

		user.Password = ""

		ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"msg":   "user created",
			"token": token,
			"user":  *user,
		})

		return nil
	}
}

func login(svc user.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user := &user.User{}
		if err := ctx.BodyParser(&user); err != nil {
			ctx.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
				"msg":   "unable to parse json",
				"error": err.Error(),
			})
			return err
		}
		user, err := svc.Login(user.Username, user.Password)
		if err != nil {
			ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"msg":   "incorrect credentials",
				"error": err.Error(),
			})
			return err
		}
		token, err := middleware.CreateToken(user.Username, user.ID)

		if err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"msg":   "unable to create token for user",
				"error": err.Error(),
			})
			return err
		}

		user.Password = ""

		ctx.Status(http.StatusOK).JSON(fiber.Map{
			"msg":   "login successful",
			"token": token,
			// "user":  *user,
		})
		return nil
	}
}

//UserEndpoints manage user endpoints
func UserEndpoints(app *fiber.App, svc user.Service) {
	app.Post("/api/user/create", register(svc))
	app.Post("/api/user/login", login(svc))

}
