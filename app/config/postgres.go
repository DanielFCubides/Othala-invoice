package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "gorm.io/driver/postgres"
	"log"
	"os"
	"othala/app/products/repositories"
)

func init() {
	err := Injector.Provide(NewPostgresConnection)
	if err != nil {
		log.Println("Error providing Postgres connection:", err)
	}
}

// getURL retrieves the URL to connection to SQL database.
func getURL(params ...string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", params[0], params[1], params[2], params[3], params[4])
}

type PostgresConnection struct {
	db *gorm.DB
}

// NewPostgresConnection retrieves a sql connection to MySQL server
func NewPostgresConnection() (Connection, error) {
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

	db, err := gorm.Open("postgres", url)
	db.LogMode(true)

	migrateModels(db)

	if err != nil {
		return nil, err
	}

	return &PostgresConnection{db: db}, nil
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&repositories.Product{})
}

func (c *PostgresConnection) GetDatabase() *gorm.DB {
	return c.db
}
