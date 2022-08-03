package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/kaium123/practice/utility"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type IProductsController interface {
	GetProduct(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetProducts(c echo.Context) error
}

type ProductsController struct {
	productRepo    repository.IProductsRepo
	productUtility utility.IProductsUtility
}

func NewProductsController(productRepo repository.IProductsRepo, productUtility utility.IProductsUtility) IProductsController {
	return &ProductsController{
		productRepo:    productRepo,
		productUtility: productUtility,
	}
}

func (p *ProductsController) Create(c echo.Context) error {

	var req = new(model.Product)
	if err := c.Bind(req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	if err := validate.Struct(req); err != nil {
		err := p.productUtility.GetErrorFeilds(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}

	err := p.productRepo.Insert(req)
	if err != nil {
		errRes := errors.ErrInternalServerErr("Data didn't create, something went wrong")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusCreated, req)
}

func (p *ProductsController) Delete(c echo.Context) error {

	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, errRes := p.ProductGetById(idx)
	if product == (model.Product{}) {
		return c.JSON(http.StatusNotFound, "That product dont exist")
	}
	if errRes != nil {
		return c.JSON(errRes.Status, errRes)
	}

	if err = p.productRepo.Delete(idx); err != nil {
		errRes := errors.ErrInternalServerErr("Data didn't deleted")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, "Data successfully deleted")
}

func (p *ProductsController) GetProduct(c echo.Context) error {
	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, errRes := p.ProductGetById(idx)

	if product == (model.Product{}) {
		return c.JSON(http.StatusNotFound, "Data not found")
	}

	if errRes != nil {
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusFound, product)
}

func (p *ProductsController) GetProducts(c echo.Context) error {
	var products []model.Product
	prefix := c.QueryParam("search")

	products, err := p.productRepo.SearchAll(prefix)

	if len(products) == 0 {
		return c.JSON(http.StatusNotFound, "Data Not Found")
	}

	if err != nil {
		errRes := errors.ErrInternalServerErr("Something went wrong")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, products)
}

func (p *ProductsController) ProductGetById(idx int) (model.Product, *errors.ErrRes) {

	product, err := p.productRepo.SearchById(idx)

	if err != nil {
		errRes := errors.ErrInternalServerErr("Something went wrong")
		return product, errRes
	}

	return product, nil
}

func (p *ProductsController) Update(c echo.Context) error {
	var req model.Product
	if err := c.Bind(&req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	errFields := p.productUtility.CheckJsonBody(req)
	if len(errFields) > 0 {
		return c.JSON(http.StatusExpectationFailed, errFields)
	}

	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, errRes := p.ProductGetById(idx)
	if product == (model.Product{}) {
		return c.JSON(http.StatusNotFound, "That data dont exist")
	}

	if errRes != nil {
		return c.JSON(errRes.Status, errRes)
	}

	product = p.productUtility.UpdateByOldData(product, req)

	if err = p.productRepo.Update(product); err != nil {
		errRes := errors.ErrInternalServerErr("Data didn't Update, some thing went wrong")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, product)
}
