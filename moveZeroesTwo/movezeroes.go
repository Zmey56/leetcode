package moveZeroesTwo

//Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

func moveZeroes(nums []int) {
	//Create a variable to keep track of the position of the last non-zero element
	lastNonZeroFoundAt := 0
	//Iterate through the array
	for i := 0; i < len(nums); i++ {
		//If the current element is not 0
		if nums[i] != 0 {
			//Swap the current element with the last non-zero element
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			//Increment the position of the last non-zero element
			lastNonZeroFoundAt++
		}
	}
}
