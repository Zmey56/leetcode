package addDigits

import "strconv"

//Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	digitCount := strconv.Itoa(num)

	for len(digitCount) > 1 {
		sum := 0
		for _, digit := range digitCount {
			digitInt, _ := strconv.Atoi(string(digit))
			sum += digitInt
		}
		digitCount = strconv.Itoa(sum)
	}

	result, _ := strconv.Atoi(digitCount)

	return result

}
