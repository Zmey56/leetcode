package main

import (
	"math"
)

func increasingTriplet(nums []int) bool {
	firstMin := math.MaxInt32
	secondMin := math.MaxInt32

	for _, num := range nums {
		if num <= firstMin {
			firstMin = num
		} else if num <= secondMin {
			secondMin = num
		} else {
			// A number greater than both firstMin and secondMin is found
			return true
		}
	}

	return false
}
