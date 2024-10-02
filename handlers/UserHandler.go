package handlers

import (
	"sales-api/models"
	"sales-api/repository"
	"sales-api/response"
	"sales-api/statusCode"
	"sales-api/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	return response.SuccessResponseWithData(c, statusCode.Ok, "Successfull", users)
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.userUsecase.GetUserById(uint(id))
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	if user.ID == 0 {
		return response.SuccessResponse(c, statusCode.NoContent, "No data found")
	}
	return response.SuccessResponseWithData(c, statusCode.Ok, "Success", user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	validationErr := user.Validate(true)
	if validationErr != nil {
		return response.ErrorResponseWithData(c, statusCode.BadRequest, "Invalid data", validationErr)
	}
	caretedUser, err := h.userUsecase.CreateUser(user)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}

	return response.SuccessResponseWithData(c, statusCode.Created, "User Created Successfully.", caretedUser)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := h.userUsecase.DeleteUser(uint(id))

	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}

	return response.SuccessResponse(c, statusCode.Ok, "User deleted successfully")
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}

	validationErr := user.Validate(false)

	if validationErr != nil {
		return response.ErrorResponseWithData(c, statusCode.BadRequest, "Invalid Data", validationErr)
	}

	updatedUser, err := h.userUsecase.UpdateUser(user)
	if err != nil {
		return response.ErrorResponse(c, statusCode.InternalServerError, err.Error())
	}
	return response.SuccessResponseWithData(c, statusCode.Ok, "User updated successfully", updatedUser)

}
