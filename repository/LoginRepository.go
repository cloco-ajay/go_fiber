package repository

import (
	"encoding/base64"
	"sales-api/config"
	"sales-api/jwt"
	"sales-api/models"
	"sales-api/utils"
	"strings"
	"time"

	"gorm.io/gorm"
)

type LoginRepository interface {
	Login(email string, password string) (interface{}, error)
	VerifyEmail(encryptedInfo string) error
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) Login(email string, password string) (interface{}, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).Where("email_verified_at IS NOT NULL").First(&user).Error
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
func (r *loginRepository) VerifyEmail(encryptedInfo string) error {
	var user models.User
	decryptedInfo, err := base64.StdEncoding.DecodeString(encryptedInfo)
	if err != nil {
		return err
	}
	data := strings.Split(string(decryptedInfo), "-")
	email := data[0]
	id := data[len(data)-1]
	now := time.Now()
	err = config.DB.Model(&user).Where("email = ? AND id = ?", email, id).Updates(models.User{EmailVerifiedAt: &now}).Error

	return err

}
