package main

import (
	"github.com/gin-gonic/gin"

	"apiorange/db/moke"
	"apiorange/service"
)

func main() {
	s := service.New(moke.New())

	r := gin.Default()
	r.GET("/ping", ping)

	r.POST("/users", s.ServiceCreateUser)
	r.GET("/users", s.ServiceGetAllUser)
	r.POST("/products", s.ServiceCreateProduct)
	r.GET("/products", s.ServiceGetAllProduct)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
