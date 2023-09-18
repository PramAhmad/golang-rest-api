package main

import (
	"belajar/api/person"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	config := cors.DefaultConfig()
	router := gin.Default()
	router.GET("/person", person.Getperson)
	router.GET("/person/:id", person.SinglePerson)
	router.PUT("/person/:id", person.UpdatePerson)
	router.POST("/person", person.PostPerson)
	config.AllowOrigins = []string{"*"}

	router.Run("127.0.0.1:8080")
}
