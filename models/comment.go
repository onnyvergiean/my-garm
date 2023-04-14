package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for comments
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Message is required"`
	UserID uint
	PhotoID uint
	User *User	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Photo *Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


type CommentCreate struct {
	Message string `json:"message"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}