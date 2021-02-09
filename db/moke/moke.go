package moke

import (
	"apiorange/db"
	"apiorange/models"
)

var _ db.Store = &Moke{}

func New() *Moke {
	return &Moke{}
}

type Moke struct {
	ListUser    []*models.User
	ListProduct []*models.Product
}

func (m *Moke) CreateUser(u *models.User) error {
	m.ListUser = append(m.ListUser, u)
	return nil
}
func (m *Moke) GetAllUser() ([]*models.User, error) {
	return m.ListUser, nil
}
func (m *Moke) CreateProduct(p *models.Product) error {
	m.ListProduct = append(m.ListProduct, p)
	return nil
}
func (m *Moke) GetAllProduct() ([]*models.Product, error) {
	return m.ListProduct, nil
}
