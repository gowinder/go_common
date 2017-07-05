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

var mysqlAddr string = "127.0.0.1"
var mysqlUser string = "root"
var mysqlPwd	string = "asdf"
var mysqlPort	int = 3306
var mysqlDb string = "test"

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
	client := &db.MysqlClient{Addr: mysqlAddr, User: mysqlUser, Pwd: mysqlPwd, Port: mysqlPort, Db: mysqlDb}
	err := client.Init(false)
	if err != nil {
		fmt.Println("testMysql init failed, ", err)
	}

	rows, err := client.Client.Query("select count(0) from testdata")
	if err != nil {
		fmt.Println("testmysql query failed, ", err)
	}

	var count int
	if err := rows.Scan(&count); err != nil {
		fmt.Println("testmysql scan row failed, ", err)
	}
	fmt.Println("testmysql scan row result: ", count)
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