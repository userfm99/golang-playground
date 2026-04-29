package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type StaticContent struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Img        string `json:"image"`
	Background string `json:"background,omitempty"`
	PromoName  string `json:"promoName,omitempty"`
	PromoDesc  string `json:"promoDesc"`
	DeepLink   string `json:"deepLink"`
}

func main() {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping, err := c.Ping().Result()

	fmt.Println(ping, err)
	var staticContents = make([]StaticContent, 0)

	content, err := c.Get("app:home:partnerDeals").Result()

	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("string json: %s", content)

	errUnmarshal := json.Unmarshal([]byte(content), &staticContents)
	if errUnmarshal != nil {
		fmt.Errorf("unmarshall error: %s", errUnmarshal.Error())
	}

	fmt.Printf("\n\nstatic contents: %v", staticContents)
}
