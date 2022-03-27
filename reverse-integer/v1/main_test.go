package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		desc     string
		input    int
		expected int
	}{
		{"Should return zero for >32 bit sized integers", 2147483648, 0},
		{"Should return zero for negative >32 bit sized integers", -2147483648, 0},
		{"Should return 321 when input is 123", 123, 321},
		{"Should return -321 when input is -123", -123, -321},
		{"Should return 21 when input is 120", 120, 21},
		{`Should return 0 when input is 1534236469 
			because its opposite is bigger than a int32 can bear`, 1534236469, 0},
	}

	for _, tc := range cases {

		r := Reverse(tc.input)
		if r != tc.expected {
			t.Fatalf("%s: expected: %d got %d", tc.desc, tc.expected, r)
		}

	}
}
