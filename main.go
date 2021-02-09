package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type User struct {
	UID      int    `json:"uid"`
	UserName string `json:"user_name"`
	Password string `json:"pass"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (u User) MarshalJSON() ([]byte, error) {
	aux := struct {
		UID      int    `json:uid`
		UserName string `json:user_name`
		Name     string `json:name`
		Email    string `json:email`
	}{
		UID:      u.UID,
		UserName: u.UserName,
		Name:     u.Name,
		Email:    u.Email,
	}

	return json.Marshal(aux)
}

type Product struct {
	UID         int      `json:"uid"`
	Title       string   `json:"title"`
	Description string   `json:"desc"`
	Reference   int      `json:"ref"`
	Images      []string `json:"imgs"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", ping)

	r.POST("/users", CreateUser)
	r.GET("/users", GetAllUser)
	r.POST("/products", CreateProduct)
	r.GET("/products", GetAllProduct)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

var (
	ListUser    []User
	ListProduct []Product
)

func GetAllUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": ListUser,
	})
}

func GetAllProduct(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": ListProduct,
	})
}

func CreateUser(ctx *gin.Context) {
	var u User
	ctx.BindJSON(&u)
	ListUser = append(ListUser, u)
	ctx.JSON(200, u)
}

func CreateProduct(ctx *gin.Context) {
	var p Product
	ctx.BindJSON(&p)
	ListProduct = append(ListProduct, p)
	ctx.JSON(200, p)
}
