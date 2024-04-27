package reverseStr

//Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.
//
//If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters,
//then reverse the first k characters and leave the other as original.

func reverseStr(s string, k int) string {
	ans := ""
	temp := ""
	j := 0
	for i := 0; i < len(s); i++ {
		if j%2 == 0 {
			temp = string(s[i]) + temp
		} else {
			temp = temp + string(s[i])
		}
		if i == (k - 1) {
			ans = ans + temp
			temp = ""
			j++
		}
		if i == ((j+1)*(k))-1 {
			ans = ans + temp
			temp = ""
			j++
		}
	}
	return ans + temp
}
