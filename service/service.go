package service

import (
	"github.com/gin-gonic/gin"

	"apiorange/db"
	"apiorange/models"
)

func New(db db.Store) *Service {
	return &Service{
		DB: db,
	}
}

type Service struct {
	DB db.Store
}

func (s *Service) ServiceGetAllUser(ctx *gin.Context) {
	us, _ := s.DB.GetAllUser()
	ctx.JSON(200, gin.H{
		"message": us,
	})
}

func (s *Service) ServiceGetAllProduct(ctx *gin.Context) {
	ps, _ := s.DB.GetAllProduct()
	ctx.JSON(200, gin.H{
		"message": ps,
	})
}

func (s *Service) ServiceCreateUser(ctx *gin.Context) {
	var u models.User
	ctx.BindJSON(&u)
	s.DB.CreateUser(&u)
	ctx.JSON(200, u)
}

func (s *Service) ServiceCreateProduct(ctx *gin.Context) {
	var p models.Product
	ctx.BindJSON(&p)
	s.DB.CreateProduct(&p)
	ctx.JSON(200, p)
}
