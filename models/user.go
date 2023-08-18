package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" valid:"required~Full Name is Required"`
	Email    string    `gorm:"not null; uniqueIndex" json:"email" valid:"required~your email is required, email - Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~password is required, minstringlength(6)~password has to have a minimum length of 6 char"`
	Products []Product `gorm:"constraint:onUpdate:CASCADE,OnDelete:SET NULL;" json:"products"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(user)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
