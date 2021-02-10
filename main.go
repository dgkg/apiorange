package main

import (
	"github.com/gin-gonic/gin"

	"apiorange/db/sqlite"
	"apiorange/service"
)

func main() {
	s := service.New(sqlite.New("moke.db"))

	r := gin.Default()

	r.POST("/users", s.ServiceCreateUser)
	r.GET("/users", s.ServiceGetAllUser)
	r.POST("/products", s.ServiceCreateProduct)
	r.GET("/products", s.ServiceGetAllProduct)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
