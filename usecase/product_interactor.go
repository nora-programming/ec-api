package usecase

import (
	"mime/multipart"

	"github.com/nora-programming/ec-api/domain"
)

type ProductInteractor struct {
	ProductRepository ProductRepository
}

func (interactor *ProductInteractor) Create(userID int, title string, description string, price int, file *multipart.FileHeader) (*domain.Product, error) {
	return interactor.ProductRepository.Create(userID, title, description, price, file)
}

func (interactor *ProductInteractor) List() ([]domain.Product, error) {
	return interactor.ProductRepository.List()
}

func (interactor *ProductInteractor) PurchasedProducts(userID string) ([]domain.PurchasedProducts, error) {
	return interactor.ProductRepository.PurchasedProducts(userID)
}
