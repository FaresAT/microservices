package main

import (
	"fmt"
	"microservices/database"
	"microservices/handlers"

	"github.com/gin-gonic/gin"
)

const (
	Port = ":5000"
)

func main() {
	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		handlers.DefaultHandler(c)
	})

	router.POST("/shorten-url", func(c *gin.Context) {
		handlers.CreateShortenedURL(c)
	})
	router.GET("/:shortenedURL", func(c *gin.Context) {
		handlers.HandleShortenedRedirect(c)
	})

	err := router.Run(Port)
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: %v", err))
	}
}
