package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//ConnectDB yields a gorm db object
func ConnectDB() (*gorm.DB, error) {

	//Heroku
	if os.Getenv("DATABASE_URL") != "" {
		return gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}
	//local
	conn, err := pq.ParseURL(os.Getenv("DB_URI"))
	fmt.Println(conn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
}
