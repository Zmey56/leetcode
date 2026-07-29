package main

// 724. Find Pivot Index

// Given an array of integers nums, calculate the pivot index of this array.

// The pivot index is the index where the sum of all the numbers strictly to the left of
//  the index is equal to the sum of all the numbers strictly to the index's right.

// If the index is on the left edge of the array, then the left sum is 0 because
//  there are no elements to the left. This also applies to the right edge of the array.

// Return the leftmost pivot index. If no such index exists, return -1.
func pivotIndex(nums []int) int {
    left, right := 0, len(nums)-1

    sumLeft, totalSum := 0, 0

    for left <= right{
        if left == right{
            totalSum += nums[left]
        }else{
            totalSum = totalSum + nums[left] + nums[right]
        }
        left++
        right--
    }

    for i, n := range nums{
        if sumLeft == totalSum - sumLeft - n{
            return i
        }
        sumLeft+=n
    }
    return -1
}