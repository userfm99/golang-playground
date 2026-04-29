package main

import (
	"log"
)

func main() {
	s := []int{1, 2, 4, 1, 5, 2}
	c := make(chan int)
	go sum(s[len(s)/2:], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c

	log.Println(x, y)
}

func sum(n []int, c chan int) {
	sum := 0
	for _, v := range n {
		sum += v
	}

	c <- sum
}
