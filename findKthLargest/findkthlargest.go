package main

import (
    "container/heap"
    "fmt"
)

// MinHeap is a custom type for a min-heap of integers.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func findKthLargest(nums []int, k int) int {
    minHeap := &MinHeap{}

    // Initialize the min-heap with the first k elements.
    for i := 0; i < k; i++ {
        heap.Push(minHeap, nums[i])
    }

    // For the remaining elements, compare them with the smallest element in the heap.
    // If an element is larger, replace the smallest element with the larger one.
    for i := k; i < len(nums); i++ {
        if nums[i] > (*minHeap)[0] {
            heap.Pop(minHeap)
            heap.Push(minHeap, nums[i])
        }
    }

    // The top element of the min-heap will be the kth largest element.
    return (*minHeap)[0]
}

func main() {
    nums1 := []int{3, 2, 1, 5, 6, 4}
    k1 := 2
    fmt.Println("Example 1:", findKthLargest(nums1, k1)) // Output: 5

    nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
    k2 := 4
    fmt.Println("Example 2:", findKthLargest(nums2, k2)) // Output: 4
}
