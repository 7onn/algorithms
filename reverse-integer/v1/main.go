package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	if x <= math.MinInt32 || x >= math.MaxInt32 {
		return 0
	}

	var negativeNumber bool = false
	if x < 0 {
		negativeNumber = true
	}

	stringNumber := fmt.Sprint(x)
	if negativeNumber {
		// Remove the minus sign
		stringNumber = stringNumber[1:]
	}
	size := len(stringNumber)

	chars := []string{}
	for _, c := range stringNumber {
		chars = append(chars, string(c))
	}

	for i := 0; i <= size-1; i++ {
		if i >= size/2 {
			continue
		}

		chars[i], chars[size-1-i] = chars[size-1-i], chars[i]
	}

	algarisms := strings.Join(chars, "")
	if negativeNumber {
		algarisms = "-" + algarisms
	}

	number, _ := strconv.Atoi(algarisms)
	if number <= math.MinInt32 || number >= math.MaxInt32 {
		return 0
	}

	return number
}

func main() {
	fmt.Println(Reverse(1534236469))
}
