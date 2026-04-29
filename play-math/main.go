package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var price float64
	var qty int

	price = 2167
	qty = 3

	fmt.Println(float64(qty))
	subtotal := math.Round(price * float64(qty))

	log.Print(subtotal)
}
