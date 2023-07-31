package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(2))    // Output: "II"
	fmt.Println(intToRoman(12))   // Output: "XII"
	fmt.Println(intToRoman(27))   // Output: "XXVII"
	fmt.Println(intToRoman(4))    // Output: "IV"
	fmt.Println(intToRoman(9))    // Output: "IX"
	fmt.Println(intToRoman(3999)) // Output: "MMMCMXCIX"
}

var romanNumerals = []struct {
	symbol string
	value  int
}{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

func intToRoman(num int) string {
	romain := strings.Builder{}
	for _, numeral := range romanNumerals {
		for num >= numeral.value {
			romain.WriteString(numeral.symbol)
			num -= numeral.value
		}
	}
	return romain.String()
}
