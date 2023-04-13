package models

type SocialMedia struct {
	GormModel
	Name string `gorm:"not null" json:"name" form:"name" valid:"required-Name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required-Social Media URL is required"`
	UserID uint
	User *User
}