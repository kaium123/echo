package utility

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func Update(product model.Product, req model.Product) model.Product {

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
func CheckErrorUpdate(req model.Product) []string {
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
				if err.Value() != 0 {
					tmp := "Invalid "
					tmp += err.StructField()
					str = append(str, tmp)
				}
			}
		}
	}

	return str
}

func CheckErrors(err error) []string {
	var str []string
	for _, err := range err.(validator.ValidationErrors) {
		s := "Invalid "
		s += err.StructField()
		str = append(str, s)
	}
	return str
}

func Check(c echo.Context) (int, interface{}, error) {
	colname := (c.QueryParam("colname"))
	if c.QueryParam("colname") == "" {
		return 0, nil, nil
	}

	if c.QueryParam("substr") == "" && c.QueryParam("prefix") == "" && c.QueryParam("suffix") == "" {
		Name, err := repository.QueryColmun(colname)
		if err != nil {
			return 1, err, err
		}
		return 1, Name, err
	}

	substr := "%" + c.QueryParam("substr") + "%"
	query := colname + " LIKE ?" //name="name LIKE ?"

	prefix := c.QueryParam("prefix") + "%"
	//query := colname + " LIKE ?" //name="name LIKE ?"

	suffix := "%" + c.QueryParam("suffix")
	//query := colname + " LIKE ?" //name="name LIKE ?"

	presuf := c.QueryParam("prefix") + "%" + c.QueryParam("suffix")

	if c.QueryParam("substr") != "" {
		Name, err := repository.SubStringSearch(query, colname, substr)
		if err != nil {
			return 1, err, err
		}
		return 1, Name, nil
	}

	if c.QueryParam("prefix") != "" && c.QueryParam("suffix") != "" {
		Name, err := repository.PreSUfSearch(query, colname, presuf)
		if err != nil {
			return 1, err, err
		}
		return 1, Name, nil
	}

	if c.QueryParam("prefix") != "" {
		Name, err := repository.SuffixStringSearch(query, colname, prefix)
		if err != nil {
			return 1, err, err
		}
		return 1, Name, nil
	}

	if c.QueryParam("suffix") != "" {
		Name, err := repository.PrefixStringSearch(query, colname, suffix)
		if err != nil {
			return 1, err, err
		}
		return 1, Name, nil
	}

	return 1, nil, nil
}
