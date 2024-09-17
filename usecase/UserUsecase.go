package usecase

import (
	"sales-api/models"
	"sales-api/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]models.User, error)
	GetUserById(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	DeleteUser(id uint) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}
type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) GetAllUsers() ([]models.User, error) {
	return uc.repo.GetAllUsers()
}

func (uc *userUsecase) GetUserById(id uint) (models.User, error) {
	return uc.repo.GetUserById(id)
}

func (uc *userUsecase) CreateUser(user models.User) (models.User, error) {
	return uc.repo.CreateUser(user)
}

func (uc *userUsecase) DeleteUser(id uint) (models.User, error) {
	return uc.repo.DeleteUser(id)
}

func (uc *userUsecase) UpdateUser(user models.User) (models.User, error) {
	return uc.repo.UpdateUser(user)
}
