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

func (r *ProductRepository) List() ([]domain.Product, error) {
	products := []domain.Product{}
	result := r.DB.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
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
		Joins("LEFT JOIN users ON users.id = purchases.buyer_id").
		Where("purchases.buyer_id = ?", userID).
		Scan(&purchasedProducts).
		Error

	if err != nil {
		return nil, err
	}

	return purchasedProducts, nil
}
