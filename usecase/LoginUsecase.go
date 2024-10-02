package usecase

import (
	"sales-api/repository"
)

type LoginUsecase interface {
	Login(email string, password string) (interface{}, error)
	VerifyEmail(encryptedInfo string) error
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
func (lu *loginUsecase) VerifyEmail(encryptedInfo string) error {
	return lu.repo.VerifyEmail(encryptedInfo)
}
