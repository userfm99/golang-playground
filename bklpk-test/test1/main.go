package main

import (
	"fmt"
	"sort"
)

func main() {
	ranks := []int{4,4,3,3,1,0}

	fmt.Println(solution(ranks))
}

func solution(ranks []int) int {
	sort.Slice(ranks, func(i, j int) bool {
		return ranks[i] > ranks[j]
	})
	fmt.Println(ranks)
	total := 0
	res := 0
	for i := 0; i < len(ranks); i++ {
		for i2 := 1; i2 < len(ranks); i2++ {
			if i > 0 && ranks[i-1] == ranks[i] {
				break
			}
			res = ranks[i] - ranks[i2]
			if res == 1 {
				total +=1
			}
		}
	}
	return total
}
