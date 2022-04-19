package repository

import (
	"mime/multipart"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO修正する
var signingKey = []byte("secret")

type UserRepository struct {
	DB *gorm.DB
}

type jwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

func (r *UserRepository) GetByID(id int) (u *domain.User, err error) {
	user := domain.User{}
	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Signin(email string, password string) (t string, u *domain.User, err error) {
	var user domain.User
	result := r.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return "", nil, result.Error
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", nil, err
	}

	claims := &jwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err = token.SignedString(signingKey)

	if err != nil {
		return t, nil, err
	}

	return t, &user, nil
}

func (r *UserRepository) Signup(email string, password string) (t string, u *domain.User, err error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	pass := string(hashed)

	user := &domain.User{
		Name:     email,
		Email:    email,
		Password: pass,
	}

	result := r.DB.Create(&user)
	if result.Error != nil {
		return "", nil, result.Error
	}

	claims := &jwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err = token.SignedString(signingKey)

	if err != nil {
		return t, nil, err
	}

	return t, user, nil
}

func (r *UserRepository) Signout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.MaxAge = 0
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	return nil
}

func (r *UserRepository) Update(userID int, name string, file *multipart.FileHeader) (*domain.User, error) {
	var user *domain.User

	// TODO fileのアップロードをする
	result := r.DB.First(&user, userID).Updates(&domain.User{Name: name})

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
