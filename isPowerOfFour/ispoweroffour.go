package isPowerOfFour

//Given an integer n, return true if it is a power of four. Otherwise, return false.
//
//An integer n is a power of four, if there exists an integer x such that n == 4x.

func isPowerOffFour(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
		if n == 1 {
			return true
		}
	}

	return false

}
