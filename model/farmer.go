package model

import "gorm.io/gorm"

type Farmer struct {
	gorm.Model
	UserId             int    `gorm:"type:varchar(100)" json:"userId"`
	User               User   `gorm:"foreignkey:UserId"`
	CompanyName        string `gorm:"type:varchar(255)" json:"companyName"`
	CompanyTradingName string `gorm:"type:varchar(255)" json:"companyTradingName"`
	CompanyEmail       string `gorm:"type:varchar(255)" json:"companyEmail"`
}
