package db

import "apiorange/models"

type Store interface {
	CreateUser(u *models.User) error
	GetAllUser() ([]*models.User, error)

	CreateProduct(p *models.Product) error
	GetAllProduct() ([]*models.Product, error)
}
