package gorm

import (
	"errors"
	"fmt"
	// "log"
	"os"
	"time"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {

	//if want to run this app to localhost, please uncomment the code below!

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	host := os.Getenv("DB_host")
	user := os.Getenv("DB_user")
	password := os.Getenv("DB_password")
	dbName := os.Getenv("DB_name")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("connect to database failed")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return gormDB, nil
}
