package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	arraySize := len(nums)
	for _, num := range nums {
		// Move only positive integers to their own index (value-1)
		for num > 0 && num-1 < arraySize {
			tmp := nums[num-1]
			nums[num-1] = num
			if tmp == num {
				break
			}
			num = tmp
		}
	}

	// Scan again and find any value-1 that doesn't match the index
	for i, num := range nums {
		if i != num-1 {
			return i + 1
		}
	}
	return arraySize + 1
}

func main() {
	// fmt.Println("Result:", firstMissingPositive([]int{-1, -3}))
	// fmt.Println("Result:", firstMissingPositive([]int{1, 1}))
	// fmt.Println("Result:", firstMissingPositive([]int{1, 2, 0}))
	fmt.Println("Result:", firstMissingPositive([]int{3, 4, -1, 1}))
	// fmt.Println("Result:", firstMissingPositive([]int{1, 200, 3, 5, 8, 7, 6, 4, 1, 2}))
	// fmt.Println("Result:", firstMissingPositive([]int{-9, 100, 1, 3, 6, 4, 1, 2}))
}
