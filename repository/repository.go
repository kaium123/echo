package repository

import (

	//"github.com/kaium123/practice/errors"

	"github.com/kaium123/practice/model"
	"gorm.io/gorm"
)

type IProductsRepo interface {
	Delete(idx int) error
	Insert(req *model.Product) error
	Search(prefix string) ([]model.Product, error)
	SearchAll(prefix string) ([]model.Product, error)
	SearchById(idx int) (model.Product, error)
	Update(product model.Product) error
}

type ProductsRepo struct {
	db *gorm.DB
}

func NewProductsRepository(DB *gorm.DB) IProductsRepo {
	return &ProductsRepo{
		db: DB,
	}
}

func (p *ProductsRepo) Delete(idx int) error {

	if err := p.db.Where("id = ?", idx).Delete(&model.Product{}); err != nil {
		return err.Error
	}

	return nil
}

func (p *ProductsRepo) Insert(req *model.Product) error {

	if err := p.db.Create(&req).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductsRepo) Search(prefix string) ([]model.Product, error) {
	var products []model.Product
	prefix += "%"
	err := p.db.Where("name LIKE ?", prefix).Find(&products)

	if err.RowsAffected == 0 {
		return products, err.Error
	}

	if err.Error != nil {
		return products, err.Error
	}

	return products, nil
}

func (p *ProductsRepo) SearchAll(prefix string) ([]model.Product, error) {

	var products []model.Product
	if prefix != "" {
		return p.Search(prefix)
	}

	err := p.db.Find(&products)
	if err != nil {
		return products, err.Error
	}

	return products, nil
}

func (p *ProductsRepo) SearchById(idx int) (model.Product, error) {

	var product model.Product

	err := p.db.First(&product, idx)

	if err.RowsAffected == 0 {
		return product, nil
	}

	if err.Error != nil {
		return product, err.Error
	}

	return product, nil
}

func (p *ProductsRepo) Update(product model.Product) error {

	if err := p.db.Save(&product); err != nil {
		return err.Error
	}

	return nil
}
