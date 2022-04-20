package repository

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nora-programming/ec-api/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) Create(userID int, title string, description string, price int, file *multipart.FileHeader) (*domain.ProductWithImg, error) {
	var imgUrl string

	product := domain.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Creater_id:  userID,
	}

	result := r.DB.Create(&product)

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
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	putParams := &s3.PutObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("products/%s", strconv.Itoa(product.ID))),
		Body:   f,
	}

	_, err = svc.PutObject(putParams)

	if err != nil {
		return nil, err
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String("ec-mall-images"),
		Key:    aws.String(fmt.Sprintf("products/%s", strconv.Itoa(product.ID))),
	})

	if err != nil {
		return nil, err
	}

	url, err := req.Presign(time.Minute * 5)

	if err != nil {
		return nil, err
	}

	imgUrl = url

	productWithImg := domain.ProductWithImg{
		Product: product,
		ImgUrl:  imgUrl,
	}

	return &productWithImg, nil
}

func (r *ProductRepository) List() ([]domain.ProductWithImg, error) {
	products := []domain.Product{}
	err := r.DB.Find(&products).Error

	if err != nil {
		return nil, err
	}

	productsWithImgs := []domain.ProductWithImg{}

	aws_access_key := os.Getenv("AWS_ACCESS_KEY")
	aws_secret_key := os.Getenv("AWS_SECRET_KEY")
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(aws_access_key, aws_secret_key, ""),
	}))

	svc := s3.New(sess, &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	for _, p := range products {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String("ec-mall-images"),
			Key:    aws.String(fmt.Sprintf("products/%s", strconv.Itoa(p.ID))),
		})

		url, _ := req.Presign(time.Minute * 5)
		productWithImg := domain.ProductWithImg{
			Product: p,
			ImgUrl:  url,
		}
		productsWithImgs = append(productsWithImgs, productWithImg)
	}

	return productsWithImgs, nil
}

func (r *ProductRepository) PurchasedProducts(userID string) ([]domain.PurchasedProducts, error) {
	purchasedProducts := []domain.PurchasedProducts{}
	err := r.DB.
		Table("purchases").
		Select([]string{
			"purchases.id",
			"products.title",
			"products.description",
			"products.price",
			"users.name",
		}).
		Joins("LEFT JOIN products ON products.id = purchases.product_id").
		Joins("LEFT JOIN users ON users.id = products.creater_id").
		Where("purchases.buyer_id = ?", userID).
		Scan(&purchasedProducts).
		Error

	if err != nil {
		return nil, err
	}

	return purchasedProducts, nil
}

func (r *ProductRepository) Delete(userID int, productID string) error {
	product := domain.Product{}
	err := r.DB.Find(&product, productID).Error

	if err != nil {
		return err
	}

	if product.Creater_id != userID {
		return errors.New("削除できません")
	}

	err = r.DB.Delete(&product).Error

	if err != nil {
		return err
	}

	return nil
}
