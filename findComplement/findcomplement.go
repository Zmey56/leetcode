package findComplement

import "log"

//The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.
//
//For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.
//Given an integer num, return its complement.

func findComplement(num int) int {
	value := []int{}
	for num > 0 {
		log.Println(num%2, " - ", num/2)
		value = append(value, num%2)
		num = num / 2
	}
	log.Println(value)
	for i := 0; i < len(value); i++ {
		if value[i] == 0 {
			value[i] = 1
		} else {
			value[i] = 0
		}
	}
	log.Println(value)

	result := 0
	for i := 0; i < len(value); i++ {
		result += value[i] * pow(2, i)
		log.Println(value[i], " - ", pow(2, i))
	}

	return result

}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
