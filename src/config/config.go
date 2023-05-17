package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// functions for connect to mysql database
func ConnectDatabase() *gorm.DB {
	checkENV := godotenv.Load()
	if checkENV != nil {
		print("Can't find .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect mysql database")
	} else {
		print("Connections success")
	}

	return db
}

// functions for disconnect from mysql database
func DisconnectDatabase(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		panic("Failed to disconnect mysql database")
	}

	database.Close()
}
