package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{"one": 1, "two": 2}
	floats := map[string]float64{"twelve": 12, "thirteen": 13}

	fmt.Printf("sum of ints: %d\nSum of floats: %f\n", SumOfMap(ints), SumOfMap(floats))
}

func SumOfMap[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}
