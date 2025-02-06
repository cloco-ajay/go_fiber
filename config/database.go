package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() map[string]interface{} {
	godotenv.Load()

	// database connection
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	fmt.Println(dbName, dbUser, dbPassword, dbHost, dbPort)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database")
		fmt.Println("Error message:", err.Error())
	}
	DB = db
	return map[string]interface{}{
		"db": db,
	}
}
