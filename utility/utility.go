package utility

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

type IProducts interface {
	UpdateByOldData(product model.Product, req model.Product) model.Product
	CheckJsonBody(req model.Product) []string
	GetErrorFeilds(err error) []string
	Search(c echo.Context) ([]string, *errors.ErrRes)
}

type products struct {
	productRepo repository.IProducts
}

func NewProductsUtility(productRepo repository.IProducts) IProducts {
	return &products{
		productRepo: productRepo,
	}
}
func (p *products) UpdateByOldData(product model.Product, req model.Product) model.Product {

	tmp := product
	product = req
	product.ID = tmp.ID

	if product.Name == "" {
		product.Name = tmp.Name
	}

	if product.Key == 0 {
		product.Key = tmp.Key
	}

	if product.Price == 0 {
		product.Price = tmp.Price
	}

	if product.Details == "" {
		product.Details = tmp.Details
	}
	return product
}
func (p *products) CheckJsonBody(req model.Product) []string {
	var str []string
	err := validate.Struct(req)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			if err.Kind() == reflect.String {
				if err.Value() != "" {
					tmp := "Invalid "
					tmp += err.StructField()
					str = append(str, tmp)
				}
			}

			if err.Kind() == reflect.Int32 {
				val := err.Value()
				var zero int32 = 0
				if val != zero {
					tmp := "Invalid "
					tmp += err.StructField()
					str = append(str, tmp)
				}
			}
		}
	}

	return str
}

func (p *products) GetErrorFeilds(err error) []string {
	var str []string
	for _, err := range err.(validator.ValidationErrors) {
		s := "Invalid "
		s += err.StructField()
		str = append(str, s)
	}
	return str
}

func (p *products) Search(c echo.Context) ([]string, *errors.ErrRes) {

	colname := "name"
	query := colname + " LIKE ?" //name="name LIKE ?"
	prefix := c.QueryParam("search") + "%"

	Name, err := p.productRepo.Search(query, prefix)
	if err != nil {
		return Name, err
	}
	return Name, nil
}
