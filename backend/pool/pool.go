package pool

import "github.com/gomodule/redigo/redis"

type Pool struct {
	*redis.Pool
}

func NewPool() {

}
