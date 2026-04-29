package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		attr, _ := e.DOM.Children().Attr("ERate-buy")
		fmt.Println(attr)

	})

	c.Visit("https://www.bca.co.id/en/informasi/kurs")
}
