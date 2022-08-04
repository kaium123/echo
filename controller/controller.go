package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type ProductsController struct {
	productRepo repository.IProductsRepo
}

func NewProductsController(productRepo repository.IProductsRepo) *ProductsController {
	return &ProductsController{
		productRepo: productRepo,
	}
}

func (p *ProductsController) Create(c echo.Context) error {

	var req = new(model.Product)
	if err := c.Bind(req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	err := p.productRepo.Insert(req)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusCreated, req)
}

func (p *ProductsController) Delete(c echo.Context) error {
	var product *model.Product
	var err error
	var idx int

	idx, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, err = p.ProductGetById(idx)
	if product == nil {
		return c.JSON(http.StatusNotFound, "That product dont exist")
	}

	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	if err = p.productRepo.Delete(idx); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, "Data successfully deleted")
}

func (p *ProductsController) GetProduct(c echo.Context) error {
	var product *model.Product
	var err error
	var idx int

	idx, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, err = p.ProductGetById(idx)

	if product == nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusFound, product)
}

func (p *ProductsController) GetProducts(c echo.Context) error {
	var products []model.Product
	prefix := c.QueryParam("search")

	products, err := p.productRepo.SearchAll(prefix)

	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, products)
}

func (p *ProductsController) ProductGetById(idx int) (*model.Product, error) {
	var product *model.Product
	var err error

	product, err = p.productRepo.SearchById(idx)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductsController) Update(c echo.Context) error {

	var product *model.Product
	var err error
	var req model.Product

	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	if err := c.Bind(&req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	product, err = p.ProductGetById(idx)
	if product == nil {
		return c.JSON(http.StatusNotFound, "That data dont exist")
	}

	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	product = UpdateByOldData(product, &req)
	if err := validate.Struct(product); err != nil {
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	if err = p.productRepo.Update(product); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateByOldData(product *model.Product, req *model.Product) *model.Product {

	if req.Name != "" {
		product.Name = req.Name
	}

	if req.Key != 0 {
		product.Key = req.Key
	}

	if req.Price != 0 {
		product.Price = req.Price
	}

	if req.Details != "" {
		product.Details = req.Details
	}
	fmt.Println(product)
	return product
}
