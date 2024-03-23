package addStrings

//Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
//
//You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
//You must also not convert the inputs to integers directly.

func addStrings(num1 string, num2 string) string {
	var res string
	carry := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		var x, y, sum int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum = x + y + carry
		if sum >= 10 {
			carry = 1
			res = 
		}

	}



	return res
}
