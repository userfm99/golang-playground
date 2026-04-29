package main

import "fmt"

func main() {
	// input := []int{4, 2, 5, 8, 7, 3, 7}
	input := []int{14, 21, 16, 35, 22}
	// input := []int{5, 5, 5, 5, 5, 5}

	fmt.Println(Solution(input))
}

func Solution(A []int) int {
	var even, odd int

	for i, num := range A {
		var sum int

		if i == len(A)-1 {
			sum = num + A[0]
		} else {
			sum = num + A[i+1]
		}

		if sum%2 == 0 {
			if i%2 == 0 {
				even++

				continue
			}

			odd++
		}
	}

	if even > odd {
		return even
	}

	if even < odd {
		return odd
	}

	return odd
}
