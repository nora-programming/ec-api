package usecase

import "github.com/nora-programming/ec-api/domain"

type PurchaseInteractor struct {
	PurchaseRepository PurchaseRepository
}

func (interactor *PurchaseInteractor) Create(userID int, productId int) error {
	return interactor.PurchaseRepository.Create(userID, productId)
}

func (interactor *PurchaseInteractor) GetSales(userID int) ([]domain.Sale, error) {
	return interactor.PurchaseRepository.GetSales(userID)
}
