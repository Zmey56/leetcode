package addStrings

import (
	"log"
	"strconv"
	"strings"
	"time"
)

//Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
//
//You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
//You must also not convert the inputs to integers directly.

func addStrings(num1 string, num2 string) string {
	//  find bigger string length
	timeNow := time.Now()
	n1, n2 := len(num1), len(num2)
	if n1 < n2 {
		num1 = strings.Repeat("0", n2-n1) + num1
	} else {
		num2 = strings.Repeat("0", n1-n2) + num2
	}

	carry := 0
	res := make([]string, len(num1)+1)

	for i := len(num1) - 1; i >= 0; i-- {
		n1, _ := strconv.Atoi(string(num1[i]))
		n2, _ := strconv.Atoi(string(num2[i]))
		sum := n1 + n2 + carry
		carry = sum / 10
		res[i+1] = strconv.Itoa(sum % 10)
		log.Println("N1 ", n1, "N2", n2, "Sum: ", sum, "Carry: ", carry, "Res: ", res)
	}

	if carry > 0 {
		res[0] = strconv.Itoa(carry)
	}

	log.Println("Time now: ", timeNow)
	return strings.Join(res, "")

}
