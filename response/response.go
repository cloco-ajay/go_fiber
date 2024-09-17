package response

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"code":    statusCode,
		"message": message,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"code":    statusCode,
		"message": message,
	})
}

func SuccessResponseWithData(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": true,
		"code":    statusCode,
		"message": message,
		"data":    data,
	})
}

func ErrorResponseWithData(c *fiber.Ctx, statusCode int, message string, errors interface{}) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"success": false,
		"code":    statusCode,
		"message": message,
		"error":   errors,
	})
}
