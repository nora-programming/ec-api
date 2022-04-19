package repository

import (
	"github.com/nora-programming/ec-api/domain"
	"gorm.io/gorm"
)

type PurchaseRepository struct {
	DB *gorm.DB
}

func (r *PurchaseRepository) Create(userID int, productId int) error {
	purchase := domain.Purchase{
		Product_id: productId,
		Buyer_id:   userID,
	}

	result := r.DB.Create(&purchase)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PurchaseRepository) GetSales(userID int) ([]domain.Sale, error) {
	sales := []domain.Sale{}

	err := r.DB.
		Table("purchases").
		Select([]string{
			"purchases.id",
			"users.name",
			"products.price",
			"products.title",
		}).
		Joins("LEFT JOIN products ON products.id = purchases.product_id").
		Joins("LEFT JOIN users ON users.id = purchases.buyer_id").
		Where("products.creater_id = ?", userID).
		Scan(&sales).
		Error

	if err != nil {
		return nil, err
	}

	return sales, nil
}
