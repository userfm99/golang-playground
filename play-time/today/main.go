package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, time.Local)
	fmt.Println(today)
}
