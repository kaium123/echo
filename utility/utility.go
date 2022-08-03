package utility

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
)

var validate = validator.New()

type IProductsRepo interface {
	UpdateByOldData(product model.Product, req model.Product) model.Product
	CheckJsonBody(req model.Product) []string
	GetErrorFeilds(err error) []string
}

type ProductsRepo struct {
	productRepo repository.IProductsRepo
}

func NewProductsUtility(productRepo repository.IProductsRepo) IProductsRepo {
	return &ProductsRepo{
		productRepo: productRepo,
	}
}
func (p *ProductsRepo) UpdateByOldData(product model.Product, req model.Product) model.Product {

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
func (p *ProductsRepo) CheckJsonBody(req model.Product) []string {
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

func (p *ProductsRepo) GetErrorFeilds(err error) []string {
	var str []string
	for _, err := range err.(validator.ValidationErrors) {
		s := "Invalid "
		s += err.StructField()
		str = append(str, s)
	}
	return str
}
