package usecase

import (
	"github.com/nora-programming/ec-api/domain"
)

type PurchaseRepository interface {
	Create(userID int, productId int) error
	GetSales(userID int) ([]domain.Sale, error)
}
