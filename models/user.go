package models

type User struct {
	GormModel
	Username string `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required-Username is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required-Email is required,email-Email is invalid"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required-Password is required,minstringlength(6)-Password must be at least 6 characters"`
	Age 	int    `gorm:"not null" json:"age" form:"age" valid:"required-Age is required,min(8)-Age must be at least 8 years old"`
}

