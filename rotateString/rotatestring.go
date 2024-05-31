package rotateString

//Given two strings s and goal, return true if and only if s can become goal after some number of shifts on s.
//
//A shift on s consists of moving the leftmost character of s to the rightmost position.
//
//For example, if s = "abcde", then it will be "bcdea" after one shift.

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		s = shift(s)
		if s == goal {
			return true
		}
	}
	return false
}

func shift(s string) string {
	r := ""
	for i := 1; i < len(s); i++ {
		r = r + string(s[i])
	}
	return r + string(s[0])
}
