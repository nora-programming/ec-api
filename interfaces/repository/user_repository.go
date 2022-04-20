package repository

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

func (r *UserRepository) GetByID(id int) (u *domain.UserWithImg, err error) {
	user := domain.User{}
	var url string

	err = r.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	aws_access_key := os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	_, err = svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
	})

	if err == nil {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
		})

		url, _ = req.Presign(time.Minute * 5)
	}

	userWithImg := domain.UserWithImg{
		User:   user,
		ImgUrl: url,
	}

	return &userWithImg, nil
}

func (r *UserRepository) Signin(email string, password string) (t string, u *domain.UserWithImg, err error) {
	var (
		url  string
		user domain.User
	)
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

	aws_access_key := os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	_, err = svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
	})

	if err == nil {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
		})

		url, _ = req.Presign(time.Minute * 5)
	}

	userWithImg := domain.UserWithImg{
		User:   user,
		ImgUrl: url,
	}

	return t, &userWithImg, nil
}

func (r *UserRepository) Signup(email string, password string) (t string, u *domain.UserWithImg, err error) {
	var (
		url string
	)
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

	aws_access_key := os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	_, err = svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
	})

	if err == nil {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
		})

		url, _ = req.Presign(time.Minute * 5)
	}

	userWithImg := domain.UserWithImg{
		User:   *user,
		ImgUrl: url,
	}

	return t, &userWithImg, nil
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

func (r *UserRepository) Update(userID int, name string, file *multipart.FileHeader) (*domain.UserWithImg, error) {
	var (
		user *domain.User
		url  string
	)

	result := r.DB.First(&user, userID).Updates(&domain.User{Name: name})

	if result.Error != nil {
		return nil, result.Error
	}

	aws_access_key := os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	if file != nil {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		putParams := &s3.PutObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
			Body:   f,
		}

		_, err = svc.PutObject(putParams)

		if err != nil {
			return nil, err
		}

	}

	_, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
	})

	if err == nil {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("users/%s", strconv.Itoa(user.ID))),
		})

		url, _ = req.Presign(time.Minute * 5)
	}

	userWithImg := domain.UserWithImg{
		User:   *user,
		ImgUrl: url,
	}

	return &userWithImg, nil
}
