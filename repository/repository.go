package repository

import (
	"fmt"

	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"gorm.io/gorm"
)

type IProducts interface {
	Insert(req *model.Product) *errors.ErrRes
	SearchById(idx int) (model.Product, *errors.ErrRes)
	Update(product model.Product) *errors.ErrRes
	Delete(idx int) *errors.ErrRes
	SearchAll() ([]model.Product, *errors.ErrRes)
	Search(query string, prefix string) ([]string, *errors.ErrRes)
}

type products struct {
	db *gorm.DB
}

func NewProductsRepository(DB *gorm.DB) IProducts {
	return &products{
		db: DB,
	}
}

func (p *products) Insert(req *model.Product) *errors.ErrRes {

	err := p.db.Create(&req).Error
	if err != nil {
		return errors.ErrInternalServerErr("Data didn't insert, something went wrong")
	}

	return nil
}

func (p *products) SearchById(idx int) (model.Product, *errors.ErrRes) {

	var product model.Product
	err := p.db.First(&product, idx)
	if err.RowsAffected == 0 {
		return product, errors.ErrNotFound("Content not found")
	}

	if err.Error != nil {
		return product, errors.ErrInternalServerErr("Something went wrong")
	}

	return product, nil
}

func (p *products) Update(product model.Product) *errors.ErrRes {

	err := p.db.Save(&product)
	if err.Error != nil {
		return errors.ErrInternalServerErr("Something went wrong")
	}

	return nil
}

func (p *products) Delete(idx int) *errors.ErrRes {

	err := p.db.Where("id = ?", idx).Delete(&model.Product{})
	if err != nil {
		return errors.ErrInternalServerErr("Something went wrong")
	}

	return nil
}

func (p *products) SearchAll() ([]model.Product, *errors.ErrRes) {

	var products []model.Product
	err := p.db.Find(&products)
	fmt.Println(products)

	fmt.Println(err.RowsAffected)
	if err.RowsAffected == 0 {
		errRes := errors.ErrNotFound("Data Not Found")
		return products, errRes
	}

	if err.Error != nil {
		errRes := errors.ErrInternalServerErr("Something went wrong")
		return products, errRes
	}

	return products, nil
}

func (p *products) Search(query string, prefix string) ([]string, *errors.ErrRes) {
	var Name []string
	err := p.db.Model(&model.Product{}).Where(query, prefix).Pluck("name", &Name)
	if err.RowsAffected == 0 {
		errRes := errors.ErrNotFound("Not Found")
		return Name, errRes
	}

	if err.Error != nil {
		errRes := errors.ErrInternalServerErr("Something went wrong")
		return Name, errRes
	}

	return Name, nil
}
