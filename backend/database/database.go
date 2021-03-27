package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hyperxpizza/url-shortener/backend/encoder"
)

type Database struct {
	client *redis.Client
}

type Item struct {
	URL  string `json:"url" redis:"url"`
	Hits int    `json:"hits" redis:"hits"`
}

func NewDatabase() Database {
	c := redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return Database{client: c}
}

func (d Database) IDexists(id uint64) bool {
	if exists := d.client.Exists(context.Background(), strconv.FormatUint(id, 10)); exists.Val() == 0 {
		return false
	}

	return true
}

func (d Database) Insert(url string, expiresAt time.Duration) (string, error) {

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	//generate id
	var id uint64
	for {
		id = rand.Uint64()
		if !d.IDexists(id) {
			break
		}
	}

	fmt.Println("Key: ", id)

	i := Item{URL: url, Hits: 0}
	json, err := json.Marshal(i)
	if err != nil {
		log.Fatalf("json.Marshal failed: %v\n", err)
		return "", err
	}

	err = d.client.Set(context.Background(), strconv.FormatUint(id, 10), json, expiresAt).Err()
	if err != nil {
		log.Fatalf("client.Set failed: %v\n", err)
		return "", err
	}

	return encoder.Encode(id), nil
}

func (d Database) Get(encodedID string) (*Item, error) {
	decodedID, err := encoder.Decode(encodedID)
	if err != nil {
		log.Fatalf("encoder.Decode failed: %v\n", err)
		return nil, err
	}

	if !d.IDexists(decodedID) {
		return nil, fmt.Errorf("Key does not exist")
	}

	s := d.client.Get(context.Background(), strconv.FormatUint(decodedID, 10))
	if s.Err() != nil {
		log.Fatal(err)
		return nil, err
	}

	var i Item
	data, err := s.Bytes()
	if err != nil {
		log.Fatalf("converting to bytes failed: %v\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &i)
	if err != nil {
		log.Fatalf("json.Unmarshal failed: %v\n", err)
		return nil, err
	}

	return &i, nil
}

func (d Database) UpdateHits(encodedID string) error {
	decodedID := encoder.Decode(encodedID)
	if err != nil {
		log.Fatalf("encoder.Decode failed: %v\n", err)
		return nil, err
	}

}
