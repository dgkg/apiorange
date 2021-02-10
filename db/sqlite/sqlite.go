// Package sqlite helps to connect with an SQLite database
//
package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"apiorange/db"
	"apiorange/models"
	"apiorange/util"
)

var _ db.Store = &SQLite{}

// SQLite is the struct that help for CRUDs.
type SQLite struct {
	db *gorm.DB
}

func New(path string) *SQLite {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{}, &models.User{})

	return &SQLite{
		db: db,
	}
}

func (s *SQLite) CreateUser(u *models.User) error {
	u.Password = util.PasswordHashing(u.Password)
	return s.db.Create(u).Error
}

func (s *SQLite) GetAllUser() ([]*models.User, error) {
	var us []*models.User
	return us, s.db.Find(&us).Error
}

func (s *SQLite) CreateProduct(p *models.Product) error {
	return s.db.Create(p).Error
}

func (s *SQLite) GetAllProduct() ([]*models.Product, error) {
	var ps []*models.Product
	return ps, s.db.Find(&ps).Error
}
