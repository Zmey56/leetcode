package KthLargest

import (
	"math"
	"sort"
)

//Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
//Implement KthLargest class:
//
//KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
//int add(int val) Appends the integer val to the stream and returns the element representing the kth largest element in the stream.

type KthLargest struct {
	K    int
	Data []int
	Cur  int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return KthLargest{
		K:    k - 1,
		Data: nums,
		Cur:  math.MinInt32,
	}
}

func (this *KthLargest) Add(val int) int {
	if val >= this.Cur {
		this.Data = append(this.Data, val)
		for j := len(this.Data) - 1; j > 0 && this.Data[j] > this.Data[j-1]; j-- {
			this.Data[j], this.Data[j-1] = this.Data[j-1], this.Data[j]
		}
		this.Cur = this.Data[this.K]

	}
	return this.Cur
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
