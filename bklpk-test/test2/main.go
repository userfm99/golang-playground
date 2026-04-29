package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution(5, 3))
}

func Solution(A int, B int) string {
	// write your code in Go 1.4

	var a = make([]string, A)
	var b = make([]string, B)

	i := 0

	for i < A {
		a[i] = "a"
		i++
	}

	i = 0

	for i < B {
		b[i] = "b"
		i++
	}

	i = 0

	fmt.Println(a, b)

	var result string
	jump := 0
	for i < (A+B) {
		if jump < len(a) {
			if (jump + 2) > len(a) {
				jump = len(a)-1
			}
			aSlice := a[jump: jump+2]
			fmt.Println(aSlice)
		}

		if jump < len(b) {
			bSlice := a[jump: jump+2]
			fmt.Println(bSlice)

		}

		i ++
		jump += 2
	}

	return result
}