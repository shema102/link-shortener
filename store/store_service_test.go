package store

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var testStoreService = &Store{}

func init() {
	envError := godotenv.Load("../.env")
	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPort = os.Getenv("REDIS_PORT")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	testStoreService = InitStore(redisHost+":"+redisPort, redisPassword)
}

func TestInitStore(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestSaveAndRetrieveShortUrl(t *testing.T) {
	initialLink := "https://www.google.com"
	shortURL := "Jsz4k57oAX"

	// Persist data mapping
	saveErr := testStoreService.SaveShortUrl(shortURL, initialLink)

	assert.Nil(t, saveErr)

	// Retrieve initial URL
	retrieved, retrievedErr := testStoreService.RetrieveLongUrl(shortURL)

	assert.Nil(t, retrievedErr)

	assert.Equal(t, initialLink, retrieved)
}
