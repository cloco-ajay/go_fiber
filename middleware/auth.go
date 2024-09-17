package middleware

import (
	"sales-api/response"
	"sales-api/statusCode"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleare(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return response.ErrorResponse(c, statusCode.Unauthorized, "Invalid or expired token")
			}
			return nil
		},
	}))
}
