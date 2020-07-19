package fizzbuzz

import "strconv"

func FizzBuzz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	}

	return strconv.Itoa(number)
}
