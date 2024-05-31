package countPrimeSetBits

//Given two integers left and right, return the count of numbers in the inclusive range [left, right] having
//a prime number of set bits in their binary representation.
//
//Recall that the number of set bits an integer has is the number of 1's present when written in binary.
//
//For example, 21 written in binary is 10101, which has 3 set bits.

func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if isPrime(countBits(i)) {
			count++
		}
	}
	return count
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func countBits(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1
	}
	return count
}
