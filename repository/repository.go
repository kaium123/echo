package repository

import (
	"fmt"

	"github.com/kaium123/practice/errors"
	"github.com/kaium123/practice/model"
	"gorm.io/gorm"
)

type IProductsRepo interface {
	Delete(idx int) error
	Insert(req *model.Product) error
	SearchAll(keyword string) ([]model.Product, error)
	FindById(idx int) (*model.Product, error)
	SearchByName(Name string) error
	Update(product *model.Product) error
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

	if err := p.db.Where("id = ?", idx).Delete(&model.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductsRepo) Insert(req *model.Product) error {

	if err := p.db.Create(&req).Error; err != nil {
		return err
	}

	return nil
}

func (p *ProductsRepo) SearchByKeword(keyword string) ([]model.Product, error) {
	var products []model.Product
	keyword = "%" + keyword + "%"
	queryExc := p.db.Where("name LIKE ?", keyword).Find(&products)

	if queryExc.RowsAffected == 0 {
		return nil, nil
	}

	if queryExc.Error != nil {
		return nil, queryExc.Error
	}

	return products, nil
}

func (p *ProductsRepo) SearchByName(Name string) error {
	var products []model.Product
	queryExc := p.db.Where("name LIKE ?", Name).Find(&products)
	fmt.Println(Name)
	fmt.Println(products)

	if queryExc.RowsAffected != 0 {
		return errors.ErrExist
	}

	if queryExc.Error != nil {
		return queryExc.Error
	}

	return nil
}

func (p *ProductsRepo) SearchAll(keyword string) ([]model.Product, error) {

	var products []model.Product
	if keyword != "" {
		return p.SearchByKeword(keyword)
	}

	queryExc := p.db.Find(&products)
	if queryExc.RowsAffected == 0 {
		return nil, nil
	}

	if queryExc.Error != nil {
		return nil, queryExc.Error
	}

	return products, nil
}

func (p *ProductsRepo) FindById(idx int) (*model.Product, error) {

	var product *model.Product

	queryExc := p.db.First(&product, idx)

	if queryExc.RowsAffected == 0 {
		return nil, errors.ErrDataNotFound
	}

	if queryExc.Error != nil {
		return nil, queryExc.Error
	}

	return product, nil
}

func (p *ProductsRepo) Update(product *model.Product) error {

	if err := p.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
