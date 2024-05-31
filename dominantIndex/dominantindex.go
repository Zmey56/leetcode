package dominantIndex

//You are given an integer array nums where the largest integer is unique.
//
//Determine whether the largest element in the array is at least twice as much as every other number in the array.
//If it is, return the index of the largest element, or return -1 otherwise.

func dominantIndex(nums []int) int {
	max := 0
	secondMax := 0
	maxIndex := 0
	for i, num := range nums {
		if num > max {
			secondMax = max
			max = num
			maxIndex = i
		} else if num > secondMax {
			secondMax = num
		}
	}
	if max >= 2*secondMax {
		return maxIndex
	}
	return -1

}
