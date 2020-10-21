package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code        string
	Title       string
	CreditCards []Review `gorm:"foreignKey:ProdId;"`
}

type Review struct {
	gorm.Model
	Asin    string
	Title   string
	Content string
	ProdId  uint
}
