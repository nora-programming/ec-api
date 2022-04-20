package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/nora-programming/ec-api/interfaces/repository"
	"github.com/nora-programming/ec-api/usecase"

	"gorm.io/gorm"
)

type ProductController struct {
	Interactor usecase.ProductInteractor
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		Interactor: usecase.ProductInteractor{
			ProductRepository: &repository.ProductRepository{
				DB: db,
			},
		},
	}
}

func (controller *ProductController) Create(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	file, _ := c.FormFile("file")
	title := c.FormValue("title")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")
	userID := c.Get("user_id").(int)
	price, _ := strconv.Atoi(priceStr)

	product, err := controller.Interactor.Create(userID, title, description, price, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, product)
}
func (controller *ProductController) List(c echo.Context) error {
	products, err := controller.Interactor.List()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, products)
}

func (controller *ProductController) PurchasedProducts(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	userIDInt, _ := c.Get("user_id").(int)
	userID := strconv.Itoa(userIDInt)
	purchased_products, err := controller.Interactor.PurchasedProducts(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, purchased_products)
}

func (controller *ProductController) Delete(c echo.Context) error {
	if c.Get("user_id") == nil {
		return nil
	}
	userIDInt, _ := c.Get("user_id").(int)
	productId := c.Param("id")

	err := controller.Interactor.Delete(userIDInt, productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusOK)
}
