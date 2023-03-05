package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shema102/link-shortener/controller"
	"github.com/shema102/link-shortener/store"
	"log"
	"os"
)

func main() {
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	storeService := store.InitStore(redisHost+":"+redisPort, redisPassword)

	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Here will be the home page",
		})
	})

	engine.POST("/shorten", func(c *gin.Context) {
		controller.CreateShortUrl(c, storeService)
	})

	engine.GET("/:shortUrl", func(c *gin.Context) {
		controller.HandleShortUrlRedirect(c, storeService)
	})

	ginPort := ":" + os.Getenv("PORT")

	ginError := engine.Run(ginPort)
	if ginError != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", ginError))
	}
}
