package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(solution(1000))

}

func solution(N int) int {
	var result int
	var result_temp int
	var calc bool

	for N > 0 {
		represent := strconv.FormatInt(int64(N), 2)
		fmt.Println(represent)
		if N%2 == 1 {
			if !calc {
				calc = true
			} else {
				if result_temp > result {
					result = result_temp
				}
				result_temp = 0
			}
		} else {
			if calc {
				result_temp++
			}
		}
		N = N / 2
	}

	return result
}
