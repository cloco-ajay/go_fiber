package routes

import (
	"sales-api/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FreeRoutes(app *fiber.App, db *gorm.DB) {

	// login handler and routes
	loginHandlers := handlers.NewLoginHandler(db)
	app.Post("/login", loginHandlers.Login)

}
