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
		{"Should return 7 when input is [4,5,6,8]", []int{4, 5, 6, 8}, 7},
		{"Should return 200 when input is [197,198,199]", []int{197, 198, 199}, 200},
	}

	for _, tc := range cases {
		r := firstMissingPositive(tc.input)
		if r != tc.expected {
			t.Fatalf("%s: expected: %d got %d", tc.desc, tc.expected, r)
		}
	}
}
