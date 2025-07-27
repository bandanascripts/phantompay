package redis

import goRedis "github.com/redis/go-redis/v9"

var GlobalRedisClient *goRedis.Client

func Connect() {

	GlobalRedisClient = goRedis.NewClient(&goRedis.Options{
		Addr: "localhost:6379",
	})
}
