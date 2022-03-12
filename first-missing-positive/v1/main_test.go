package main

import (
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		desc     string
		input    []int
		expected int
	}{
		{"Should return 1 where there is only negative numbers", []int{-3, -2, -1}, 1},
		{"Should return 1 when input is [4,5,6,8]", []int{4, 5, 6, 8}, 1},
		{"Should return 2 when input is [1,4,5,6,8]", []int{1, 4, 5, 6, 8}, 2},
		{"Should return 2 when input is [3,4,-1,1]", []int{3, 4, -1, 1}, 2},
		{"Should return 2 when input is [1,1]", []int{1, 1}, 2},
		{"Should return 3 when input is [1,1,4]", []int{1, 2, 0}, 3},
		{"Should return 5 when input is [1,2,3,4,197,198,199]", []int{1, 2, 3, 4, 197, 198, 199}, 5},
	}

	for _, tc := range cases {
		r := firstMissingPositive(tc.input)
		if r != tc.expected {
			t.Fatalf("%s: expected: %d got %d", tc.desc, tc.expected, r)
		}
	}
}
