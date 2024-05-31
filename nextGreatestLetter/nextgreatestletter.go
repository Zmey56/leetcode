package nextGreatestLetter

//You are given an array of characters letters that is sorted in non-decreasing order, and a character target.
//There are at least two different characters in letters.
//
//Return the smallest character in letters that is lexicographically greater than target.
//If such a character does not exist, return the first character in letters.

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < len(letters) {
		return letters[left]
	}

	return letters[0]
}
