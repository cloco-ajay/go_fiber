package routes

import (
	"sales-api/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FreeRoutes(app *fiber.App, db *gorm.DB) {

	// login handler and routes
	loginHandlers := handlers.NewLoginHandler(db)
	userHandlers := handlers.NewUserHandler(db)
	
	// user handlers and routes
	app.Post("/users/create", userHandlers.CreateUser)

	
	app.Post("/login", loginHandlers.Login)
	app.Get("/verify-email/:encryptedInfo", loginHandlers.VerifyEmail).Name("VerifyEmail")


}
