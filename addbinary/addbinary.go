//67. Add Binary
//Given two binary strings a and b, return their sum as a binary string.
//
//Example 1:
//
//Input: a = "11", b = "1"
//Output: "100"
//Example 2:
//
//Input: a = "1010", b = "1011"
//Output: "10101"
//
//
//Constraints:
//
//1 <= a.length, b.length <= 104
//a and b consist only of '0' or '1' characters.
//Each string does not contain leading zeros except for the zero itself.

package main

import (
	"log"
	"strconv"
)

func addBinary(a string, b string) string {
	result := ""
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		log.Println("i0", i)
		log.Println("j0", j)
		log.Println("carry0", carry)
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			log.Println("a[i] - '0'", a[i]-'0')
			i--
		}

		if j >= 0 {
			sum += int(b[j] - '0')
			log.Println("b[j] - '0'", b[j]-'0')
			j--
		}

		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
		log.Println("SUM", sum)
		log.Println("i", i)
		log.Println("j", j)
		log.Println("carry", carry)
	}

	return result
}

func main() {

	a := "11"
	b := "1"
	log.Println(addBinary(a, b))

	//a1 := "1010"
	//b1 := "1011"
	//log.Println(addBinary(a1, b1))
}
