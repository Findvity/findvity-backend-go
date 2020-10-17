package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BRO3886/findvity-backend/pkg/event"
	"github.com/BRO3886/findvity-backend/pkg/group"
	"github.com/BRO3886/findvity-backend/pkg/user"
	"github.com/BRO3886/findvity-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	//set initial viper config
	// viper.SetConfigName(".env")
	// viper.SetConfigType("env")
	// viper.AddConfigPath(".")

	//read .env

	if os.Getenv("ON_SERVER") != "True" {
		// Loading the .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	//connect db and attach logger
	db, err := utils.ConnectDB()
	if err != nil {
		log.Panicln(fmt.Errorf("Error Opening Database %s", err))
	}

	//perform migrations
	db.AutoMigrate(&user.User{}, &event.Event{}, &group.Group{})
	log.Println("connected to db")

	//close db connection
	// defer db.Close()

	//start fiber
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "India/Kolkata",
		Output:     os.Stdout,
	}))

	//healthcheck
	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Status(http.StatusOK).JSON(fiber.Map{
			"status":  "ok",
			"message": "operational",
		})
		return nil
	})

	utils.MakeUserHandlers(app, db)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
