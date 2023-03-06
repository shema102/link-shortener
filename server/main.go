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
	env := os.Getenv("ENV")

	if env != "production" {
		envError := godotenv.Load()
		if envError != nil {
			log.Fatal("Error loading .env file")
		}
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	storeService := store.InitStore(redisHost+":"+redisPort, redisPassword)

	engine := gin.Default()

	engine.StaticFile("/", "../client/dist/index.html")
	engine.StaticFile("/favicon.svg", "../client/dist/favicon.svg")
	engine.Static("/assets", "../client/dist/assets")

	engine.POST("/api/shorten", func(c *gin.Context) {
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
