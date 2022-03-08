package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for ki, vi := range nums {
		for kj, vj := range nums {
			if ki == kj {
				continue
			}
			if vi+vj == target {
				return []int{ki, kj}

			}
			continue
		}
		continue
	}
	return []int{-1}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
