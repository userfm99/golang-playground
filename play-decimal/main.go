package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var x float64
	x = 84573
	fmt.Println(math.Floor(x*100) / 100) // 12.34 (round down)
	fmt.Println(math.Round(x*100) / 100) // 12.35 (round to nearest)
	fmt.Println(math.Ceil(x*100) / 100)  // // 12.35 (round up)

	s := fmt.Sprintf("%.2f", math.Round(x*100)/100)
	m, _ := strconv.ParseFloat(s, 64)
	fmt.Println(m)

	fmt.Println(round2DecimalPlaces(x))

	xm := 0.0000001
	if xm < 1 && xm > 0 {
		fmt.Println("got")
	}
}

func round2DecimalPlaces(amount float64) float64 {
	s := fmt.Sprintf("%.2f", math.Round(amount*100)/100)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return math.Round(amount)
	}
	return f
}
