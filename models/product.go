package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UID         int     `json:"uid"`
	Title       string  `json:"title"`
	Description string  `json:"desc"`
	Reference   int     `json:"ref"`
	Images      []Image `json:"imgs" gorm:"foreignKey:ProductRefer"`
}

type Image struct {
	gorm.Model
	URI          string `json:"uri"`
	ProductRefer uint
}
