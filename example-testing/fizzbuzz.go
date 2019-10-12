package demo

import (
	"strconv"
)

func fizzbuzz(number int) string {
	if number%5 == 0 && number%3 == 0 {
		return "fizzbuzz"
	}
	if number%3 == 0 {
		return "fizz"
	}
	if number%5 == 0 {
		return "buzz"
	}

	return strconv.Itoa(number)
}
