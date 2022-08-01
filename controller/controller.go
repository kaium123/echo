package controller

import (
	"fmt"
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

type IProducts interface {
	GetProduct(c echo.Context) error
	Post(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Get(c echo.Context) error
}

type products struct {
	productRepo    repository.IProducts
	productUtility utility.IProducts
}

func NewProductsController(productRepo repository.IProducts, productUtility utility.IProducts) IProducts {
	return &products{
		productRepo:    productRepo,
		productUtility: productUtility,
	}
}
func (p *products) GetProduct(c echo.Context) error {
	product, err := p.ProductGetById(c)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusFound, product)
}

func (p *products) Post(c echo.Context) error {

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
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, req)
}

func (p *products) ProductGetById(c echo.Context) (model.Product, *errors.ErrRes) {

	var product model.Product
	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := errors.ErrBadRequest(err.Error())
		return product, &errRes
	}

	product, errRes := p.productRepo.SearchById(idx)
	if errRes != nil {
		return product, errRes
	}

	return product, nil
}

func (p *products) Update(c echo.Context) error {
	product, err := p.ProductGetById(c)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	var req model.Product
	if err := c.Bind(&req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	errFields := p.productUtility.CheckJsonBody(req)
	if len(errFields) > 0 {
		return c.JSON(http.StatusExpectationFailed, errFields)
	}

	product = p.productUtility.UpdateByOldData(product, req)
	Err := p.productRepo.Update(product)
	if Err != nil {
		return c.JSON(Err.Status, Err)
	}

	return c.JSON(http.StatusOK, product)
}

func (p *products) Delete(c echo.Context) error {

	_, errRes := p.ProductGetById(c)
	if errRes != nil {
		return c.JSON(errRes.Status, errRes)
	}
	idx, err := strconv.Atoi(c.Param("id"))
	//idx, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	errRes = p.productRepo.Delete(idx)
	if err != nil {
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, " successfully deleted")
}

func (p *products) Get(c echo.Context) error {
	var products []model.Product
	if c.QueryParam("search") != "" {
		Name, err := p.productUtility.Search(c)

		if err != nil {
			return c.JSON(err.Status, err)
		}
		return c.JSON(http.StatusOK, Name)
	}

	products, errRes := p.productRepo.SearchAll()
	fmt.Println(errRes)
	if errRes != nil {
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, products)
}
