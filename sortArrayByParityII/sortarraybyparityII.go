package sortArrayByParityII

//Given an array of integers nums, half of the integers in nums are odd, and the other half are even.
//
//Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.
//
//Return any answer array that satisfies this condition.

func sortArrayByParityII(nums []int) []int {
	oddSlice := []int{}
	evenSlice := []int{}
	for _, v := range nums {
		if v%2 == 0 {
			oddSlice = append(oddSlice, v)
		} else {
			evenSlice = append(evenSlice, v)
		}
	}

	var res []int
	for i := 0; i < len(nums); i++ {
		if len(oddSlice) == 0 || len(evenSlice) == 0 {
			if len(oddSlice) > 0 {
				res = append(res, oddSlice...)
			}
			if len(evenSlice) > 0 {
				res = append(res, evenSlice...)
			}
			break
		}
		if i%2 == 0 {
			res = append(res, oddSlice[len(oddSlice)-1])
			oddSlice = oddSlice[:len(oddSlice)-1]
		} else {
			res = append(res, evenSlice[len(evenSlice)-1])
			evenSlice = evenSlice[:len(evenSlice)-1]
		}
	}

	return res
}
