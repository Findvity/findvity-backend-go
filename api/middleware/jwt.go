package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/BRO3886/findvity-backend/api"
	"github.com/gofiber/fiber/v2"

	"github.com/dgrijalva/jwt-go"
)

//Token struct
type Token struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

//BasicJWTAuth auth token checker
func BasicJWTAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		tokenHeader := string(ctx.Request().Header.Peek("Authorization"))

		if tokenHeader == "" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": api.ErrTokenMissing.Error(),
			})

		}
		tk := &Token{}
		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtsecret")), nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(http.StatusForbidden).JSON(fiber.Map{
				"message": api.ErrInvalidToken.Error(),
			})
		}
		return ctx.Next()
	}
}

//GetUIDFromToken to get the token from handler
func GetUIDFromToken(ctx *fiber.Ctx) string {
	tokenHeader := string(ctx.Request().Header.Peek("Authorization"))
	tk := &Token{}
	_, err := jwt.ParseWithClaims(tokenHeader, tk, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("jwtsecret")), nil
	})
	if err != nil {
		return ""
	}
	return tk.Username
}

//CreateToken used to create JWT
func CreateToken(id, username string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["ID"] = id
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Hour * 23).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("jwtsecret")))
	if err != nil {
		return "", err
	}
	return token, nil
}
