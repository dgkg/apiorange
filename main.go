package main

import "github.com/gin-gonic/gin"

type User struct {
	UID      int    `json:uid`
	UserName string `json:user_name`
	Password string `json:pass`
	Name     string `json:name`
	Email    string `json:email`
}

type Product struct {
	UID         int      `json:uid`
	Title       string   `json:title`
	Description string   `json:desc`
	Reference   int      `json:ref`
	Images      []string `json:imgs`
}

func main() {

	r := gin.Default()
	r.GET("/ping", ping)

	r.POST("/users", CreateUser)
	//r.GET("/users", GetAllUser)
	r.POST("/products", CreateProduct)
	//r.GET("/products", GetAllProduct)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

var (
	listUser    []User
	listProduct []Product
)

func CreateUser(ctx *gin.Context) {
	var u User
	ctx.BinJSON(&u)
	listUser = append(listUser, u)
	ctx.JSON(200, u)
}

func CreateProduct(ctx *gin.Context) {
	var p Product
	ctx.BinJSON(&p)
	listProduct = append(listProduct, p)
	ctx.JSON(200, p)
}
