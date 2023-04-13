package models


type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required-Message is required"`
	UserID uint
	PhotoID uint
	User *User
	Photo *Photo
}