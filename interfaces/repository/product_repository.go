package repository

import (
	"mime/multipart"

	"github.com/nora-programming/ec-api/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (r *ProductRepository) Create(userID int, title string, description string, price int, file *multipart.FileHeader) (*domain.Product, error) {
	// TODO fileをアップロードする
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

	return &product, nil
}
