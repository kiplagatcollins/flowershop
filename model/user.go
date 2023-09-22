package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(100)" json:"username"`
	FirstName   string `gorm:"type:varchar(255)" json:"firstName"`
	LastName    string `gorm:"type:varchar(255)" json:"lastName"`
	Email       string `gorm:"type:varchar(255)" json:"email"`
	Address     string `gorm:"type:varchar(255)" json:"address"`
	PhoneNumber string `gorm:"type:varchar(255)" json:"phoneNumber"`
	DOB         string `gorm:"type:varchar(255)" json:"dob"`
	Role        string `gorm:"type:varchar(255)" json:"role"`
}
