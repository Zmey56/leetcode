package findErrorNums

//You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error,
//one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.
//
//You are given an integer array nums representing the data status of this set after the error.
//
//Find the number that occurs twice and the number that is missing and return them in the form of an array.

func findErrorNums(nums []int) []int {
	var res []int
	myMap := make(map[int]int)
	for _, val := range nums {
		myMap[val]++
	}
	for key, val := range myMap {
		if val > 1 {
			res = append(res, key)
			break
		}
	}
	val := 0
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				val = 0
				break
			} else {
				val = i
			}
		}
		if val != 0 {
			break
		}
	}
	res = append(res, val)
	return res
}
