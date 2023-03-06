package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shema102/link-shortener/shortener"
	"github.com/shema102/link-shortener/store"
	"net/http"
	"os"
)

type ShortUrlRequest struct {
	Url    string `json:"url" binding:"required"`
	UserId string `json:"userId" binding:"required"`
}

func CreateShortUrl(c *gin.Context, store *store.Store) {
	var shortUrlRequest ShortUrlRequest

	if err := c.ShouldBindJSON(&shortUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl, err := shortener.GenerateShortUrl(shortUrlRequest.Url, shortUrlRequest.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	storeErr := store.SaveShortUrl(shortUrl, shortUrlRequest.Url)

	if storeErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": storeErr.Error()})
		return
	}

	host := os.Getenv("HOST")

	c.JSON(http.StatusOK, gin.H{"shortUrl": host + "/" + shortUrl})
}

func HandleShortUrlRedirect(c *gin.Context, store *store.Store) {
	shortUrl := c.Param("shortUrl")

	url, err := store.RetrieveLongUrl(shortUrl)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, url)
}
