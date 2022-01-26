package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	positives := []int{}
	for _, current := range A {
		if current > 0 {
			positives = append(positives, current)
		}
	}
	if len(positives) < 1 {
		return 1
	}

	sort.Ints(positives)

	size := len(positives)
	if size == 1 {
		return positives[0:1][0] + 1
	}

	var max int
	for k, current := range positives {
		if k == 0 {
			continue
		}

		previous := positives[k-1 : k][0]
		if current == previous {
			continue
		}

		if current-previous > 1 {
			return previous + 1
		}

		max = current
		continue
	}

	return max + 1
}

func main() {
	fmt.Println("Result:", Solution([]int{-1, -3}))
	fmt.Println("Result:", Solution([]int{1, 200, 3, 5, 8, 7, 6, 4, 1, 2}))
	fmt.Println("Result:", Solution([]int{-9, 100, 1, 3, 6, 4, 1, 2}))
}
