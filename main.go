//  gowinder@hotmail.com 2017/7/5 9:22
package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gowinder/go_common/db"
)

var redisAddr string = "127.0.0.1:6379"
var redisPwd string = "asdf"
var redisDb int = 0

func main() {
	fmt.Println("start svpn cache gate  version 0.1.3 ...")


	testRedis()

	err := db.GlobalRedisClientPool.Init(redisAddr, redisPwd, redisDb, 50, true, true)
	if err != nil {
		fmt.Println("GlobalRedisClientPool.Init error ", err)
		return
	}


	testMysql()

}
func testMysql() {

}

func testRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPwd, // no password set
		DB:       redisDb,  // use default DB
	})

	_, err := client.Ping().Result()
	if err == nil{
		fmt.Println("test redis", redisAddr, "ok")
	}else {
		fmt.Println("test redis", redisAddr, "failed, err:", err)
	}

	err = client.Set("testkey", "mother fucker", 0).Err()
	if err != nil {
		panic(err)
	}

}