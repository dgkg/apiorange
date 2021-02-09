package main

import "github.com/gin-gonic/gin"


type User struct {
	Username string
	Password string 
	Uid int
	Name string 
	Emmail string 
}

type Product struct {
	Title string
	Description string 
	Reference int 
	Images []string
}

func main() {
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
