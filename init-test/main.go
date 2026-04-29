package main

import (
	"log"

	"github.com/userfm99/golang-playground/init-test/initsetup"
)

func main() {
	log.Println("message from main")
	initsetup.Construct()
}
