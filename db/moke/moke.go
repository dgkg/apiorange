package moke

import (
	"apiorange/db"
	"apiorange/models"
	"apiorange/util"
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
	u.Password = util.PasswordHashing(u.Password)
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
