package database

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Database struct {
	*redis.Pool
}

func NewDatabase() Database {
	pool := redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf(":%d", 6379))
		},
	}

	return Database{&pool}
}
