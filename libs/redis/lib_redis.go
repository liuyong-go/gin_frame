package libs

import (
	"gin_frame/config"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.LoadConfig().Redis.Addr,
		Password: "",                           // no password set
		DB:       config.LoadConfig().Redis.Db, // use default DB
	})
}
