package repository

import (
	"github.com/kaium123/practice/database"
	"github.com/kaium123/practice/model"
)

var db = database.GetDB()

func Insert(req *model.Product) error {

	err := db.Create(&req).Error
	return err
}

func SearchById(idx int) (model.Product, error) {

	var tmp model.Product
	err := db.First(&tmp, idx).Error
	return tmp, err
}

func Update(product model.Product) error {

	err := db.Save(&product).Error
	return err
}

func Delete(idx int) error {

	err := db.Where("id = ?", idx).Delete(&model.Product{}).Error
	return err
}

func SearchAll() ([]model.Product, error) {

	var product []model.Product
	err := db.Find(&product).Error
	return product, err
}

func QueryColmun(colname string) (interface{}, error) {
	var Name []string
	err := db.Model(&model.Product{}).Pluck(colname, &Name).Error
	return Name, err
}

func SubStringSearch(query string, colname string, substr string) (interface{}, error) {
	var Name []string
	err := db.Model(&model.Product{}).Where(query, substr).Pluck(colname, &Name).Error
	return Name, err
}

func PreSUfSearch(query string, colname string, presuf string) (interface{}, error) {
	var Name []string
	err := db.Model(&model.Product{}).Where(query, presuf).Pluck(colname, &Name).Error
	return Name, err
}

func SuffixStringSearch(query string, colname string, suffix string) (interface{}, error) {
	var Name []string
	err := db.Model(&model.Product{}).Where(query, suffix).Pluck(colname, &Name).Error
	return Name, err
}

func PrefixStringSearch(query string, colname string, prefix string) (interface{}, error) {
	var Name []string
	err := db.Model(&model.Product{}).Where(query, prefix).Pluck(colname, &Name).Error
	return Name, err
}
