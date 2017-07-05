//  gowinder@hotmail.com 2017/7/5 9:05
package db

import (
	"github.com/go-redis/redis"
	"fmt"
	"sync"
)

type RedisClient struct{
	Client 	*redis.Client
	Addr	string
	Pwd		string
	Db		int

	Pool	*RedisClientPool
}

func (self *RedisClient) Init(ping bool) error {
	self.Client = redis.NewClient(&redis.Options{
		Addr:     self.Addr,
		Password: self.Pwd, // no password set
		DB:       self.Db,  // use default DB
	})

	if ping{
		_, err := self.Client.Ping().Result()
		if err == nil{
			fmt.Println("RedisClient.Init ping", self.Addr, "ok")
		}else {
			fmt.Println("RedisClient.Init ping", self.Addr, "failed, err:", err)
			return err
		}
	}

	return nil
}

func (self *RedisClient) ReturnToPool(){
	if self.Pool != nil{
		self.Pool.ReturnClient(self)
	}
}


type RedisClientPool struct{
	sync.Mutex
	pool	 []*RedisClient
	//	poolInUse	[]*RedisClient
	intUse	int
}

var GlobalRedisClientPool	RedisClientPool

func (self *RedisClientPool) Init(addr string, pwd string, db int, cap int, pingTest bool, breadIfError bool) error{
	self.Lock()
	defer self.Unlock()
	self.pool = make([]*RedisClient, cap)
	fmt.Println("RedisClientPool.Init begin, capacity is", cap)

	for i := 0; i < cap; i++{
		self.pool[i] = &RedisClient{Addr:addr,Pwd:pwd,Db:db}
		redisClient := self.pool[i]
		redisClient.Pool = self

		err := redisClient.Init(pingTest)
		if err != nil && breadIfError{
			return err
		}
	}


	return nil
}

func (self *RedisClientPool) GetClient() *RedisClient{
	self.Lock()
	defer self.Unlock()

	if len(self.pool) < 1{
		fmt.Println("RedisClientPool.GetClient no free client")
		return nil
	}

	index := len(self.pool) - 1
	client := self.pool[index]
	self.pool = self.pool[:index]

	//self.poolInUse = append(self.poolInUse, client)

	self.intUse += 1

	return client
}

func (self *RedisClientPool) ReturnClient(redisClient *RedisClient) {
	self.Lock()
	defer self.Unlock()


	self.pool = append(self.pool, redisClient)

	self.intUse -= 1
}
