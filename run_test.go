package test

import (
	libs_redis "gin_frame/libs/redis"
	"testing"
	"time"
)

//go test -v -run TestRedis run_test.go
func TestRedis(t *testing.T) {
	var rdb = libs_redis.Rdb
	var expire_time = time.Duration(1000) * time.Second
	rdb.Set("name", "liuyong", expire_time)
	value := rdb.Get("name")
	t.Log("test redis value")
	t.Log(value)
}
func TestDb(t *testing.T) {
	t.Log("test db")
}
