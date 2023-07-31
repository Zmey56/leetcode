package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := 123
	fmt.Println(reverse(x))

	x = -123
	fmt.Println(reverse(x))

	x = 120
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var roundedNumber int
	if x < 0 {
		roundedNumber = x * -1
	} else {
		roundedNumber = x
	}

	str := strconv.Itoa(int(roundedNumber))
	result := []string{}

	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, string(str[i]))
	}

	resultInt, _ := strconv.Atoi(strings.Join(result, ""))

	if resultInt > -2147483648 && resultInt < 2147483647 {
		if x < 0 {
			return resultInt * -1
		} else {
			return resultInt
		}
	} else {
		return 0
	}
}
