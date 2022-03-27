package main

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
)

func Reverse(x int) int {
	limit := new(big.Int).Exp(big.NewInt(2), big.NewInt(31), nil)
	if x <= int(limit.Int64())*-1 || x >= int(limit.Int64()) {
		return 0
	}
	stringNumber := fmt.Sprint(x)

	negativeNumber := false
	if x < 0 {
		negativeNumber = true
	}

	size := 0
	if negativeNumber {
		// Ignore the minus sign
		size = len(stringNumber) - 1
	} else {
		size = len(stringNumber)
	}

	chars := []string{}
	for _, c := range stringNumber {
		if string(c) == "-" {
			continue
		}
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

	number, err := strconv.Atoi(algarisms)
	if err != nil {
		log.Fatalf("string -> int parse Error\n%+v", err)
	}

	if number <= int(limit.Int64())*-1 || number >= int(limit.Int64()) {
		return 0
	}

	return number
}

func main() {
	fmt.Println(Reverse(123))
	fmt.Println(Reverse(-123))
	fmt.Println(Reverse(120))
	fmt.Println(Reverse(-2147483647))
	fmt.Println(Reverse(2147483647))
	fmt.Println(Reverse(1534236469))
}
