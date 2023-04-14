package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for social media
type SocialMedia struct {
	GormModel
	Name string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social Media URL is required"`
	UserID uint
	User *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type SocialMediaCreate struct {
	Name string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}