package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/url"
)

func main() {
	url1, err := url.Parse("http://localhost:1881")
	if err != nil {
		log.Fatal(err)
	}

	/*	url2, err := url.Parse("http://localhost:1882")
		if err != nil {
			log.Fatal(err)
		}*/

	e := echo.New()
	g := e.Group("/v3/home")
	g.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: url1,
			},
		}),
		Rewrite: map[string]string{"/v3/home/*": "/v2/home/$1"},
	}))
	e.Start(":1880")
}
