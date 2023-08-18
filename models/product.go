package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title  string `json:"title" form:"title" valid:"required ~ title is required"`
	UserID uint
	User   *User
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(product)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (product Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(product)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
