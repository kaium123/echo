package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name    string `json:"name" validate:"required,min=4,max=20"`
	Key     int32  `json:"key" validate:"required,min=4,max=20"`
	Price   int32  `json:"price" validate:"required,min=4,max=100"`
	Details string `json:"details" validate:"required,min=4,max=20"`
}
