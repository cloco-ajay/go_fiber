package handlers

import (
	"fmt"
	"sales-api/response"
	"sales-api/repository"
	"sales-api/statusCode"
	"sales-api/usecase"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginHandler struct {
	useCase usecase.LoginUsecase
}

func NewLoginHandler(db *gorm.DB) *LoginHandler {
	repo := repository.NewLoginRepository(db)
	useCase := usecase.NewLoginUsecase(repo)

	return &LoginHandler{useCase: useCase}
}

func (h *LoginHandler) Login(c *fiber.Ctx) error {

	var payload map[string]interface{}
	err := c.BodyParser(&payload)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	email := fmt.Sprint(payload["email"])
	password := fmt.Sprint(payload["password"])
	data, err := h.useCase.Login(email, password)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	return response.SuccessResponseWithData(c, statusCode.Ok, "Logged In successfully.", data)
}
