package hasAlternatingBits

import "strconv"

//Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

func hasAlternatingBits(n int) bool {
	bitsFormat := strconv.FormatInt(int64(n), 2)
	for i := 0; i < len(bitsFormat)-1; i++ {
		if bitsFormat[i] == bitsFormat[i+1] {
			return false
		}
	}
	return true
}
