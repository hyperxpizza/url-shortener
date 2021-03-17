package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return Database{client: c}
}

func (d Database) IDexists(id uint64) bool {
	res := d.client.Exists(context.Background(), strconv.FormatUint(id, 10))
	if res.Val() == 0 {
		return false
	}

	return true
}

func (d Database) Insert(url string, expiresAt time.Duration) (string, error) {
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

func (d Database) Load(encodedID string) (string, error) {
	decodedID, err := encoder.Decode(encodedID)
	if err != nil {
		log.Fatalf("encoder.Decode failed: %v\n", err)
		return "", err
	}

	s := d.client.Get(context.Background(), strconv.FormatUint(decodedID, 10))

	return s.String(), nil
}

func (d Database) GetInfo(encodedID string) error {
	decodedID, err := encoder.Decode(encodedID)
	if err != nil {
		log.Fatalf("encode.Decode failed: %v\n", err)
		return err
	}

	if !d.CheckIfKeyExists(decodedID) {
		return fmt.Errorf("this id does not exists")
	}

	//val := d.client.Get(context.Background())
	return nil
}

func (d Database) CheckIfKeyExists(decodedID uint64) bool {
	if exists := d.client.Exists(context.Background(), strconv.FormatUint(decodedID, 10)); exists.Val() == 0 {
		return false
	}

	return true
}
