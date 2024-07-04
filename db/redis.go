package db

import (
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func init() {
	opt, err := redis.ParseURL("rediss://default:AbegAAIncDFhMmQyMzg1MTZhNjQ0MzFiOWEwODZkZTNjOTRjMmI0ZHAxNDcwMDg@usable-raccoon-47008.upstash.io:6379")
	if err != nil {
		panic(err)
	}
	Redis = redis.NewClient(opt)
}
