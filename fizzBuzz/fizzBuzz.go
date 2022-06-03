package fizzBuzz

import "strconv"

func FizzBuzz(number int) (result string) {
	if number%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
