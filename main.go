package main

import (
	"fmt"
	"sales-api/config"
	"sales-api/database"
	"sales-api/middleware"
	"sales-api/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// fiber instance
	app := fiber.New()

	response := config.DatabaseConnection()
	db, ok := response["db"].(*gorm.DB)
	if !ok {
		fmt.Println("Not type gorm db")
	}
	// handling auto migration of the models/Schema
	database.AutoMigration(db)

	// routes
	routes.FreeRoutes(app, db)

	// jwt middleware
	middleware.JwtMiddleare(app)
	routes.AuthRoutes(app, db)

	app.Listen(":7000")

}
