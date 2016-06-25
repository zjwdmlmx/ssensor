package global

import "gopkg.in/redis.v3"

var R *redis.Client

func initRedis() {
	R = redis.NewClient(&redis.Options{
		Addr:     "localhost:10000",
		PoolSize: 5,
	})
}
