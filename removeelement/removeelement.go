package main

import "log"

func removeElement(nums []int, val int) int {
	k := 0 // index to track the number of elements without val
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i] // move the element to the front
			k++
		}
	}
	return k

}

func main() {
	nums1 := []int{3, 2, 2, 3}
	val1 := 3

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2

	log.Println(removeElement(nums1, val1))
	log.Println(removeElement(nums2, val2))
}
