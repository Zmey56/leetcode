package sortArrayByParity

// Given an integer array nums, move all the even integers at the beginning of the array followed by all the odd integers.
//
// Return any array that satisfies this condition.
func sortArrayByParity(nums []int) []int {
	evenSlice := []int{}
	noEvenSlice := []int{}
	for _, num := range nums {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		} else {
			noEvenSlice = append(noEvenSlice, num)
		}
	}

	return append(evenSlice, noEvenSlice...)
}

func sortArrayByParityVer2(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 > nums[right]%2 {
			nums[left], nums[right] = nums[right], nums[left]
		}

		if nums[left]%2 == 0 {
			left++
		}

		if nums[right]%2 == 1 {
			right--
		}
	}
	return nums
}
