package libs

import (
	"fmt"
	"gin_frame/config"

	"github.com/go-redis/redis"
)

var Rdb *redis.Client

func init() {
	fmt.Print(config.LoadConfig().Redis.Password)
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.LoadConfig().Redis.Addr,
		Password: config.LoadConfig().Redis.Password, // no password set
		DB:       config.LoadConfig().Redis.Db,       // use default DB
	})
}
