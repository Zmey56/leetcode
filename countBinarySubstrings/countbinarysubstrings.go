package countBinarySubstrings

//Given a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's,
//and all the 0's and all the 1's in these substrings are grouped consecutively.
//
//Substrings that occur multiple times are counted the number of times they occur.

func countBinarySubstrings(s string) int {
	countOne := 0
	countZero := 0
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if i == 0 || s[i-1] == '1' {
				countZero = 0
			}
			countZero++
			if countZero <= countOne {
				count++
			}
		} else {
			if i == 0 || s[i-1] == '0' {
				countOne = 0
			}
			countOne++
			if countOne <= countZero {
				count++
			}
		}
	}
	return count
}
