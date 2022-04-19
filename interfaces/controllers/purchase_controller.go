package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/interfaces/repository"
	"github.com/nora-programming/ec-api/usecase"

	"gorm.io/gorm"
)

type PurchaseController struct {
	Interactor usecase.PurchaseInteractor
}

func NewPurchaseController(db *gorm.DB) *PurchaseController {
	return &PurchaseController{
		Interactor: usecase.PurchaseInteractor{
			PurchaseRepository: &repository.PurchaseRepository{
				DB: db,
			},
		},
	}
}

func (controller *PurchaseController) Create(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	productId, _ := strconv.Atoi(c.FormValue("product_id"))
	userID := c.Get("user_id").(int)

	err := controller.Interactor.Create(userID, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}

func (controller *PurchaseController) GetSales(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	userID := c.Get("user_id").(int)

	sales, err := controller.Interactor.GetSales(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, sales)
}
