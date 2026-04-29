package main

import (
	"fmt"
	"io"
)

type devNull int

func main() {
	var discard io.Writer = devNull(1)
	fmt.Println(discard)
}

func (devNull) Write(p []byte) (int, error) {
	return len(p), nil
}
