package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := "42"
	fmt.Println(myAtoi(s))

	s = "   -42"
	fmt.Println(myAtoi(s))

	s = "4193 with words"
	fmt.Println(myAtoi(s))

	s = "words and 987"
	fmt.Println(myAtoi(s))

	s = "3.14159"
	fmt.Println(myAtoi(s))

	s = ""
	fmt.Println(myAtoi(s))

	s = "+1"
	fmt.Println(myAtoi(s))

	s = " "
	fmt.Println(myAtoi(s))

	s = "  "
	fmt.Println(myAtoi(s))
}

func myAtoi(s string) int {
	trimStr := strings.TrimSpace(s)

	if len(trimStr) < 1 {
		return 0
	}

	firstLetter := trimStr[0]
	positive := true
	if string(firstLetter) == "-" {
		positive = false
		trimStr = trimStr[1:]
	} else if string(firstLetter) == "+" {
		trimStr = trimStr[1:]
	}

	result := []string{}
	for m, i := range trimStr {
		if m == 0 && !unicode.IsDigit(i) {
			return 0
		}

		if unicode.IsDigit(i) {
			result = append(result, string(i))
		} else {
			break
		}
	}

	resultInt, _ := strconv.ParseInt(strings.Join(result, ""), 10, 64)

	if !positive {
		resultInt = resultInt * -1
	}

	if resultInt < -2147483648 {
		return -2147483648
	} else if resultInt > 2147483647 {
		return 2147483647
	} else {
		return int(resultInt)
	}

}
