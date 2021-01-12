package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
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
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", params[0], params[1], params[2], params[3], params[4])
}

type PostgressConnection struct {
	db *gorm.DB
}

// NewPostgressConnection retrieves a sql connection to MySQL server
func NewPostgressConnection() (Connection, error) {
	dbUsername := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_USER_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	url := getURL(dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	log.Println(url)

	db, err := gorm.Open("postgres", url)
	db.LogMode(true)

	if err != nil {
		return nil, err
	}

	return &PostgressConnection{db: db}, nil
}

func (c *PostgressConnection) GetDatabase() *gorm.DB {
	return c.db
}
