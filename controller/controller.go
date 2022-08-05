package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/kaium123/practice/types"
	"github.com/kaium123/practice/utility"

	//"github.com/kaium123/practice/utility"
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

	var req = new(types.ProductInfo)
	var model *model.Product

	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	if err := validate.Struct(req); err != nil {
		message := utility.ValidationMessage(err)
		fmt.Println(message)
		return c.JSON(http.StatusUnprocessableEntity, message)
	}

	productName := req.Name
	err := p.productRepo.SearchByName(productName)

	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	if err := utility.StructToStruct(req, &model); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	err = p.productRepo.Insert(model)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, req)
}

func (p *ProductsController) Delete(c echo.Context) error {
	//var product *model.Product
	var err error
	var idx int

	idx, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	_, err = p.productRepo.FindById(idx)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	if err = p.productRepo.Delete(idx); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, utility.DeleteMessage("Data successfully deleted"))
}

func (p *ProductsController) GetProduct(c echo.Context) error {
	var product *types.ProductInfo
	var err error
	var idx int

	idx, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	product, err = p.ProductGetById(idx)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, product)
}

func (p *ProductsController) GetProducts(c echo.Context) error {
	var products = make([]types.ProductInfo, 0)
	var model []model.Product
	keyword := c.QueryParam("search")

	model, err := p.productRepo.SearchAll(keyword)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	if err := utility.StructToStruct(model, &products); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, products)
}

func (p *ProductsController) ProductGetById(idx int) (*types.ProductInfo, error) {
	var product *types.ProductInfo
	var model *model.Product
	var err error

	model, err = p.productRepo.FindById(idx)
	if err != nil {
		return nil, err
	}

	if err := utility.StructToStruct(model, &product); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductsController) Update(c echo.Context) error {

	var product *types.ProductInfo
	var err error
	var req types.ProductInfo
	var model *model.Product

	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, "Invalid Id")
	}

	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		errRes := errors.ErrBadRequest("Invalid json body")
		return c.JSON(errRes.Status, errRes)
	}

	product, err = p.ProductGetById(idx)
	if err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	product = UpdateByOldData(product, &req)
	if err := validate.Struct(product); err != nil {
		message := utility.ValidationMessage(err)
		return c.JSON(http.StatusUnprocessableEntity, message)
	}

	if err := utility.StructToStruct(product, &model); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	if err = p.productRepo.Update(model); err != nil {
		errRes := errors.CheckErr(err, "Product")
		return c.JSON(errRes.Status, errRes)
	}

	return c.JSON(http.StatusOK, product)
}

func UpdateByOldData(product *types.ProductInfo, req *types.ProductInfo) *types.ProductInfo {

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
