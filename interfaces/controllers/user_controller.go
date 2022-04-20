package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/domain"
	"github.com/nora-programming/ec-api/interfaces/repository"
	"github.com/nora-programming/ec-api/usecase"

	"gorm.io/gorm"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

// TODO S3のimgURLを返す
type UserResponse struct {
	ID    int
	Name  string
	Email string
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				DB: db,
			},
		},
	}
}

func (controller *UserController) Me(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	user, err := controller.Interactor.Me(c.Get("user_id").(int))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Signin(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	t, u, err := controller.Interactor.Signin(user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, u)
}

func (controller *UserController) Signup(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	t, u, err := controller.Interactor.Signup(user.Email, user.Password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, u)
}

func (controller *UserController) Signout(c echo.Context) error {
	err := controller.Interactor.Signout(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, nil)
}

func (controller *UserController) Update(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	file, _ := c.FormFile("file")
	name := c.FormValue("name")
	userID := c.Get("user_id").(int)

	user, err := controller.Interactor.Update(userID, name, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}
