package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx = &gin.Context{}
)

func SetupRdb() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rdb.Ping(ctx)

	val, err := rdb.Incr(ctx, "pageviews").Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Pageviews: %d\n", val) // Pageviews: 1 (then 2, 3, ...)
}

func GetRdb() *redis.Client {
	return rdb
}
