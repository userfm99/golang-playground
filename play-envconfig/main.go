package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Environment struct
type Environment struct {
	Name string `envconfig:"NAME" required:"true"`
	Free string `envconfig:"FREE"`
}

func main() {
	var env Environment
	err := envconfig.Process("prod", &env)
	if err != nil {
		panic(err)
	}
	fmt.Printf("name of the service: %v", env.Name)
}
