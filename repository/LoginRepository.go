package repository

import (
	"sales-api/config"
	"sales-api/jwt"
	"sales-api/models"
	"sales-api/utils"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Login(email string, password string) (interface{}, error)
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) Login(email string, password string) (interface{}, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if user.ID != 0 {
		correctPassword, passErr := utils.CheckPassword(user.Password, password)
		if correctPassword == false {
			return nil, passErr
		}
		token := jwt.GenerateToken(*user.Email, int(user.ID))
		return token, err

	}
	return nil, err

}
