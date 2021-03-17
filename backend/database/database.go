package database

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type Database struct {
	client *redis.Client
}

type Item struct {
	ID        uint64 `json:"id" redis:"id"`
	URL       string `json:"url" redis:"url"`
	ExpiresAt string `json:"expires_at" redis:"expires_at"`
	Hits      int    `json:"hits" redis:"hits"`
}

func NewDatabase() Database {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return Database{client: c}
}

func (d Database) existsID(id uint64) {
	val := d.client.Get(context.Background(), strconv.FormatUint(id, 10))
	fmt.Println(val.String())
}

func (d Database) insertItem(url string, expiresAt time.Time) {
	id := rand.Uint64()
	shortLink := Item{
		ID:        id,
		URL:       url,
		ExpiresAt: expiresAt.Format("2006-01-02 15:04:05.728046 +0300 EEST"),
		Hits:      0,
	}

	err := d.client.Set()
}
