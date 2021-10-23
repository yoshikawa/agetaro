package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	// MySQL driver.
	_ "github.com/go-sql-driver/mysql"
)

// GetGormConn sets up Gorm connection.
func GetGormConn() (*gorm.DB, error) {
	err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return gorm.Open("mysql", fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?%s&%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_CHARSET"), os.Getenv("DB_PARSETIME")))
}

// GetCloudSQLConn sets up Gorm connection.
func GetCloudSQLConn() (*gorm.DB, error) {
	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	return gorm.Open("mysql", fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), socketDir, os.Getenv("DB_INSTANCE"), os.Getenv("DB_PROTOCOL")))
}
