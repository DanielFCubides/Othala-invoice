package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func init() {
	err := Injector.Provide(NewPostgressConnection)
	if err != nil {
		log.Println("Error providing Postgress connection:", err)
	}
}

// getURL retrieves the URL to connection to SQL database.
func getURL(params ...string) string {

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Bogota", params[0], params[1], params[2], params[3], params[4])
}

type PostgressConnection struct {
	db *gorm.DB
}

// NewPostgressConnection retrieves a sql connection to postgres server
func NewPostgressConnection() (Connection, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := getURL(
		dbHost,
		dbUsername,
		dbPassword,
		dbName,
		dbPort,
	)
	log.Println(url)

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		return nil, err
	}

	return &PostgressConnection{db: db}, nil
}

func (c *PostgressConnection) GetDatabase() *gorm.DB {
	return c.db
}
