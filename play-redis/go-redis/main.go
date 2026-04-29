package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	i := client.Scan(0, "other:plu:*", 0).Iterator()
	for i.Next() {
		log.Println(i.Val())
		if err := client.Del(i.Val()).Err(); err != nil {
			log.Println("err del key: ", i.Val())
		}
	}
	if err := i.Err(); err != nil {
		panic(err)
	}
}
