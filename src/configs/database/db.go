package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error){

	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env error")
	}

	host := os.Getenv("DB_Host")
	user := os.Getenv("DB_user")
	password := os.Getenv("DB_password")
	dbName := os.Getenv("DB_name")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	gormDB, err := gorm.Open(postgres.Open(config), &gorm.Config{})


	db, err := gormDB.DB()
	if err != nil {
		return nil, errors.New("connect to database failed")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return gormDB, nil
}