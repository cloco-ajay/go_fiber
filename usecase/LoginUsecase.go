package usecase

import (
	"sales-api/repository"
)

type LoginUsecase interface {
	Login(email string, password string) (interface{}, error)
}

type loginUsecase struct {
	repo repository.LoginRepository
}

func NewLoginUsecase(repo repository.LoginRepository) LoginUsecase {
	return &loginUsecase{repo: repo}
}

func (lu *loginUsecase) Login(email string, password string) (interface{}, error) {
	return lu.repo.Login(email, password)
}
