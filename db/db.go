package db

import "apiorange/models"

type Store interface {
	CreateUser(u *models.User) error
	GetAllUser() ([]models.User, error)

	CreateProduct(p *models.Porduct) error
	GetAllProduct() ([]models.Product, error)
}
