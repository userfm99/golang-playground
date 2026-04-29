package main

import "fmt"

func main() {
	adder1 := func(a, b int) int {
		return a * b
	}

	adder2 := func(a, b int) int {
		return a + b
	}

	// returned Adder type
	adderResult1 := adderition(adder1)
	adderResult2 := adderition(adder2)

	fmt.Println(adderResult1(1, 2))
	fmt.Println(adderResult2(1, 2))
}

type adder func(a, b int) int

func adderition(adder1 adder) adder {
	return func(a, b int) int {
		return 10*a + b
	}
}
