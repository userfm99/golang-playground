package main

import (
	"fmt"

	"github.com/userfm99/golang-playground/play-with-interface/animal"
	"github.com/userfm99/golang-playground/play-with-interface/circus"
)

func main() {
	var dog animal.Dog
	fmt.Println(circus.Perform(dog))
}
