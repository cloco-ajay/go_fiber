package routes

import (
	"sales-api/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {

	// user handlers and routes
	userHandlers := handlers.NewUserHandler(db)
	

	// grouping routes with prefix users
	users := app.Group("/users")
	users.Get("/", userHandlers.GetAllUsers)
	users.Get("/:id", userHandlers.GetUserById)
	users.Delete("/delete/:id", userHandlers.DeleteUser)
	users.Post("/update", userHandlers.UpdateUser)
}
