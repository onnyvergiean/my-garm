package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)


type Photo struct {
	GormModel
	Title string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption string `gorm:"not null" json:"caption" form:"caption"`
	PhotoURL string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required"`
	UserID uint 
	User *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}