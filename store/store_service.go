package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Store struct {
	redisClient *redis.Client
}

var (
	store = &Store{}
	ctx   = context.Background()
)

const CacheDuration = 12 * time.Hour

func InitStore(redisUrl string, redisPassword string) *Store {
	fmt.Printf("Connecting to Redis: %v", redisUrl)

	store.redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPassword,
		DB:       0,
	})

	pong, err := store.redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis - Error: %v", err))
	}

	fmt.Printf("Connected to Redis succesfully: %v", pong)

	return store
}

func (s *Store) SaveShortUrl(shortUrl string, longUrl string) error {
	err := s.redisClient.Set(ctx, shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) RetrieveLongUrl(shortUrl string) (string, error) {
	val, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
