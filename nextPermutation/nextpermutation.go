package nextPermutation

//A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
//
//For example, for arr = [1,2,3], the following are all the permutations of arr:
//[1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
//The next permutation of an array of integers is the next lexicographically greater permutation of its integer.
//More formally, if all the permutations of the array are sorted in one container according to their lexicographical
//order, then the next permutation of that array is the permutation that follows it in the sorted container.
//If such arrangement is not possible, the array must be rearranged as the lowest possible order (i.e., sorted in ascending order).
//
//For example, the next permutation of arr = [1,2,3] is [1,3,2].
//Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
//While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical larger rearrangement.
//Given an array of integers nums, find the next permutation of nums.
//
//The replacement must be in place and use only constant extra memory.

func nextPermutation(num []int) {
	var i, j int
	for i = len(num) - 2; i >= 0; i-- {
		if num[i] < num[i+1] {
			for j = len(num) - 1; j > i; j-- {
				if num[j] > num[i] {
					break
				}
			}
			num[i], num[j] = num[j], num[i]
			reverse(num[i+1:])
			return
		}
	}
	reverse(num)

}

func reverse(num []int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}
