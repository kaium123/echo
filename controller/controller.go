package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/kaium123/practice/model"
	"github.com/kaium123/practice/repository"
	"github.com/kaium123/practice/utility"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func Get(c echo.Context) error {
	product, _, err := UserGetById(c)
	if err != nil {
		return c.JSON(http.StatusOK, "That product not found")
	}

	return c.JSON(http.StatusOK, product)
}

func Post(c echo.Context) error {
	var req = new(model.Product)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusExpectationFailed, "There is an error")
	}

	if err := validate.Struct(req); err != nil {
		err := utility.CheckErrors(err)
		return c.JSON(http.StatusExpectationFailed, err)
	}

	err := repository.Insert(req)
	if err != nil {
		return c.JSON(http.StatusOK, "Opps data didn't insert")
	}

	return c.JSON(http.StatusOK, req)
}

func UserGetById(c echo.Context) (model.Product, int, error) {

	var tmp model.Product
	idx, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return tmp, idx, err
	}

	tmp, err = repository.SearchById(idx)
	if err != nil {
		return tmp, idx, err
	}

	return tmp, idx, nil
}

func Update(c echo.Context) error {
	product, _, err := UserGetById(c)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "That data not found")
	}

	var req model.Product
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusExpectationFailed, "failed to Bind data")
	}

	str := utility.CheckErrorUpdate(req)
	if len(str) > 0 {
		return c.JSON(http.StatusExpectationFailed, str)
	}

	product = utility.Update(product, req)
	err = repository.Update(product)
	if err != nil {
		return c.JSON(http.StatusExpectationFailed, "failed")
	}

	return c.JSON(http.StatusOK, product)
}

func Delete(c echo.Context) error {

	_, idx, err := UserGetById(c)
	if err != nil {
		c.JSON(http.StatusOK, "not found1")
	}

	err = repository.Delete(idx)
	if err != nil {
		return c.JSON(http.StatusOK, "failed")
	}

	return c.JSON(http.StatusOK, " successfully deleted")
}

func Getall(c echo.Context) error {
	var product []model.Product
	//fmt.Println(c.QueryParam("colname"))
	IsParameterPresent, result, err := utility.Check(c)
	fmt.Println(result, IsParameterPresent)
	if err != nil {
		return c.JSON(http.StatusOK, "failed")
	}
	if IsParameterPresent == 1 {
		return c.JSON(http.StatusOK, result)
	}
	product, err = repository.SearchAll()
	if err != nil {
		return c.JSON(http.StatusOK, "failed")
	}

	return c.JSON(http.StatusOK, product)
}
