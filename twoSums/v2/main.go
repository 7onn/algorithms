package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if _, exists := hashmap[complement]; exists {
			return []int{hashmap[complement], i}
		}
		hashmap[nums[i]] = i
	}
	return []int{-1}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
