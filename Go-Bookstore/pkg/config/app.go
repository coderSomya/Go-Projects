package config

import (
	// "hello/go/pkg/mod/github.com/jinzhu/gorm@v1.9.16"

	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var(
	db *gorm.DB
)

func Connect(){
	dbUsername := os.Getenv("DB_USERNAME")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

	connectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", connectionString)
    if err != nil {
    panic("Failed to connect to database")
    }

	db=d
}

func GetDB() *gorm.DB{
	return db
}
