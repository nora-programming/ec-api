package usecase

import (
	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Me(id int) (user *domain.User, err error) {
	return interactor.UserRepository.GetByID(id)
}

func (interactor *UserInteractor) Signin(email string, password string) (t string, u *domain.User, err error) {
	return interactor.UserRepository.Signin(email, password)
}

func (interactor *UserInteractor) Signup(email string, password string) (t string, u *domain.User, err error) {
	return interactor.UserRepository.Signup(email, password)
}

func (interactor *UserInteractor) Signout(c echo.Context) error {
	return interactor.UserRepository.Signout(c)
}
