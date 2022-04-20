package usecase

import (
	"mime/multipart"

	"github.com/nora-programming/ec-api/domain"
)

type ProductRepository interface {
	Create(userID int, title string, description string, price int, file *multipart.FileHeader) (*domain.ProductWithImg, error)
	List() ([]domain.ProductWithImg, error)
	PurchasedProducts(userID string) ([]domain.PurchasedProducts, error)
	Delete(userID int, productID string) error
}
