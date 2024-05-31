package selfDividingNumbers

//A self-dividing number is a number that is divisible by every digit it contains.
//
//For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.
//A self-dividing number is not allowed to contain the digit zero.
//
//Given two integers left and right, return a list of all the self-dividing numbers in the range [left, right].

func selfDividingNumbers(left int, right int) []int {
	var result []int
	for i := left; i <= right; i++ {
		if isSelfDividingNumber(i) {
			result = append(result, i)
		}
	}
	return result
}

func isSelfDividingNumber(n int) bool {
	for i := n; i > 0; i /= 10 {
		if i%10 == 0 || n%(i%10) != 0 {
			return false
		}
	}
	return true
}
