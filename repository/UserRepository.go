package repository

import (
	"sales-api/models"
	"sales-api/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetUserById(id uint) (models.User, error) {
	var user models.User
	err := r.db.Find(&user, id).Error
	return user, err
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {

	hashedPassword, hashError := utils.HashPassword(user.Password)
	if hashError != nil {
		return user, hashError
	}
	user.Password = hashedPassword

	err := r.db.Create(&user).Error

	return user, err
}

func (r *userRepository) DeleteUser(id uint) (models.User, error) {
	var user models.User

	err := r.db.Delete(&user, id).Error

	return user, err

}

func (r *userRepository) UpdateUser(user models.User) (models.User, error) {

	err := r.db.Model(&user).Updates(map[string]interface{}{"name": user.Name}).Error

	return user, err
}
