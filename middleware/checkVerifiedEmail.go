package middleware

import (
	"github.com/gofiber/fiber/v2"

)

func CheckVerifiedEmail(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})
}