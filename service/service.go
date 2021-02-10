package service

import (
	"github.com/gin-gonic/gin"

	"apiorange/db"
	"apiorange/models"
	"apiorange/util"
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

type Auth struct {
	Login string `json:"login"`
	Password string `json:"pass"`
}

func (s *Service) ServiceLogin(ctx *gin.Context) {
	var a Auth
	ctx.BindJSON(&a)
	user, err := s.DB.GetPassByLogin(a.Login)
	if err != nil {
		ctx.JSON(500, err)
		return
	} 
	if user.Password != util.PasswordHashing(a.Password) {
		ctx.JSON(403, "unknown user")
		return 
	}
	token, err := util.NewJWT("macle", a.Login)
	if err != nil {
		ctx.JSON(403, "unable to create token")
		return 
	}
		
	ctx.JSON(200, gin.H{"jwt": token})
		
}


