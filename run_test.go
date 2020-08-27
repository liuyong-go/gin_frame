package test

import (
	"context"
	"fmt"
	libs_elastic "gin_frame/libs/elastic"
	libs_etcd "gin_frame/libs/etcd"
	libs_redis "gin_frame/libs/redis"
	libs_short "gin_frame/libs/shortUrl"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"go.etcd.io/etcd/clientv3"
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
func TestEtcdGet(t *testing.T) {
	var timeout = time.Duration(10) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	var cli = libs_etcd.EtcdCli
	cli.Put(ctx, "sample_key", "sample_value")
	var value = libs_etcd.GetOne("sample_key11")
	if value == "" {
		fmt.Print("weikong")
	}
	fmt.Print(value)
	cancel()
}
func TestEtcdLease(t *testing.T) {
	var timeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var cli = libs_etcd.EtcdCli
	resp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		t.Log(err)
	}
	_, err = cli.Put(context.TODO(), "fo22o", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		t.Log(err)
	}
	time.Sleep(time.Duration(3) * time.Second)

	fmt.Println("ttl:", resp.TTL)
	// to renew the lease only once
	ka, kaerr := cli.KeepAliveOnce(context.TODO(), resp.ID)
	if kaerr != nil {
		t.Log(kaerr)
	}
	time.Sleep(time.Duration(6) * time.Second)
	fmt.Println("ttl:", ka.TTL)
	value, err := cli.Get(ctx, "fo22o")
	if err != nil {
		t.Log(err)
	}
	t.Log(value)
	for _, ev := range value.Kvs {
		t.Log(string(ev.Value))
	}

}
func TestGetGoroutine(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("error")
			fmt.Print(err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(err)
	}
	fmt.Print("go id", id)
}
func TestLock(t *testing.T) {
	rs := libs_etcd.GetLock("liuyong", 5, 3)
	if rs == true {
		fmt.Print("加锁成功")
	} else {
		fmt.Print("加锁失败")
	}
	time.Sleep(20 * time.Second)
}

//go test -v -run TestElastic run_test.go
func TestElastic(t *testing.T) {
	//var els = libs_elastic.Elas
	var con = map[string]interface{}{
		"title":   "maptest",
		"company": "soyoung",
		"content": "soyoung test 2",
	}
	// var content = `{"title":"test","name":"liuyong","content":"test hello elasticsearch"}`
	libs_elastic.CreateDocument("test", con)
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": "test",
			},
		},
	}
	libs_elastic.Search(query, "test")
}
func TestShortUrl(t *testing.T) {
	var url = "https://synewad.xinyangwang.net/v2/syshare/product"
	var shortURL = libs_short.CreateShortURL(url)
	fmt.Print("surl:" + shortURL + "end")

}
