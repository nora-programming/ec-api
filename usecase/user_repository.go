package usecase

import (
	"mime/multipart"

	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/domain"
)

type UserRepository interface {
	GetByID(id int) (user *domain.User, err error)
	Signin(email string, password string) (t string, u *domain.User, err error)
	Signup(email string, password string) (t string, u *domain.User, err error)
	Signout(c echo.Context) error
	Update(userID int, name string, file *multipart.FileHeader) (*domain.User, error)
}
