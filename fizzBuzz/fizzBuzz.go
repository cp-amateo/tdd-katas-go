package fizzBuzz

import "strconv"

func FizzBuzz(number int) (result string) {
	if number%3 == 0 {
		return "Fizz"
	}
	if number%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(number)
}
