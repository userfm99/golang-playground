package main

import (
	"log"
	"strconv"
)

func main() {
	// name := "John Doe"
	age := 21
	hello(&age)
	log.Println(age)
}

func hello(dest ...interface{}) {
	switch dtype := dest[0].(type) {
	case *string:
		log.Println("string", dtype)
		break
	case *int:
		log.Println("Integer", dtype)
		*dtype = 1
		break
	}
}

func ScanFunc(productID *uint64, term *string, handler func(json string)) {
	c := GetRedisClient()
	key := "pdp:" + strconv.FormatUint(*productID, 10) + ":" + *term
	v, err := c.Get("")
}
