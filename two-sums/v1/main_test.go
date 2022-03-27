package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		desc     string
		values   []int
		target   int
		expected []int
	}{
		{"Should return [-1] when no pair of numbers sums up to target", []int{1, 1, 1}, 3, []int{-1}},
		{"Should return [0,1] when first two numbers sums up to target", []int{1, 2, 3}, 3, []int{0, 1}},
		{"Should return [1,2] when second and third numbers sums up to target", []int{1, 2, 3}, 5, []int{1, 2}},
	}
	for _, tc := range cases {
		actual := twoSum(tc.values, tc.target)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Fatalf("%s: expected: %d got: %d for a: %d and b %d",
				tc.desc, tc.expected, actual, tc.values, tc.target)
		}
	}
}
