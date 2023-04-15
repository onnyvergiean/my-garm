package models

import (
	"errors"
	"my-garm/helpers"

	_ "my-garm/docs"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// User represents the model for users
type User struct {
	GormModel `json:"-"`
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Age 	int    `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type UserRegister struct{
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age int `json:"age"`
}

type UserLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil  {
		err = errCreate
		return
	}

	if u.Age <=8 {
		err  = errors.New("age must be greater than 8 years old")
		return 
	}

	u.Password = helpers.HashPassword(u.Password)
	err = nil
	return
}