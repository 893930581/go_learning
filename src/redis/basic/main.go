package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initClient() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr :	"localhost:6379",
		Password:	"",
		DB:	0,
	})

	_, err = redisdb.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Print("链接成功")
	return nil
}

func redisExample() {
	err := redisdb.Set("score",100,0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}
	val ,err := redisdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score",val)
	val2, err := redisdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}

func main() {
	initClient()
	redisExample()
}