package usecase

import (
	"mime/multipart"

	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Me(id int) (user *domain.UserWithImg, err error) {
	return interactor.UserRepository.GetByID(id)
}

func (interactor *UserInteractor) Signin(email string, password string) (t string, u *domain.UserWithImg, err error) {
	return interactor.UserRepository.Signin(email, password)
}

func (interactor *UserInteractor) Signup(email string, password string) (t string, u *domain.UserWithImg, err error) {
	return interactor.UserRepository.Signup(email, password)
}

func (interactor *UserInteractor) Signout(c echo.Context) error {
	return interactor.UserRepository.Signout(c)
}

func (interactor *UserInteractor) Update(userID int, name string, file *multipart.FileHeader) (*domain.UserWithImg, error) {
	return interactor.UserRepository.Update(userID, name, file)
}
