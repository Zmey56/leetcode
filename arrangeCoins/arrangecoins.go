package arrangeCoins

//You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where
//the ith row has exactly i coins. The last row of the staircase may be incomplete.
//
//Given the integer n, return the number of complete rows of the staircase you will build.

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	count := 0
	value := 1
	for n > 0 {
		count++
		n -= value
		if n < 0 {
			count--
			break
		}
		value++
		//log.Println("Count: ", count, "Value: ", value, "N: ", n)
	}
	//log.Println("RESULT Count: ", count, "Value: ", value, "N: ", n)
	return count
}
