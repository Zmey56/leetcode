package minSubArrayLen

import (
	"math"
	"sync"
)

//Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose
//sum is greater than or equal to target. If there is no such subarray, return 0 instead./
//
//
//Example 1:
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem constraint.
//Example 2:
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//Example 3:
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//Constraints:
//
//1 <= target <= 109
//1 <= nums.length <= 105
//1 <= nums[i] <= 104

func minSubArrayLen(target int, nums []int) int {
	result := 0
	sum := 0
	for i, num := range nums {
		if num >= target {
			return 1
		}
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				length := j - i + 1
				if result == 0 || length < result {
					result = length
				}
				break
			}
		}
		sum = 0
	}
	return result
}

func minSubArrayLenVer2(target int, nums []int) int {
	left := 0
	sum := 0
	result := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		// Shrink window from left while sum >= target
		for sum >= target {
			length := right - left + 1
			if result == 0 || length < result {
				result = length
			}
			sum -= nums[left]
			left++
		}
	}

	return result
}

func minSubArrayLenV3(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Determine number of goroutines based on array size
	minWorkers := min(4, n)
	if n < 100 {
		// For small arrays, sequential is faster
		return minSubArrayLenSequential(target, nums)
	}

	chunkSize := (n + minWorkers - 1) / minWorkers

	// Channel to collect results from workers
	resultCh := make(chan int, minWorkers)
	var wg sync.WaitGroup
	// Launch goroutines to process chunks
	for i := 0; i < minWorkers; i++ {
		wg.Add(1)
		start := i * chunkSize
		end := min((i+1)*chunkSize, n)

		go func(startIdx, endIdx int) {
			defer wg.Done()
			minLen := math.MaxInt32

			// Try all subarrays starting from this chunk
			for left := startIdx; left < endIdx; left++ {
				sum := 0
				for right := left; right < n; right++ {
					sum += nums[right]
					if sum >= target {
						length := right - left + 1
						if length < minLen {
							minLen = length
						}
						break
					}
				}
			}
			resultCh <- minLen
		}(start, end)
	}

	// Close channel after all goroutines finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect minimum from all workers
	result := math.MinInt32
	for localMin := range resultCh {
		result = min(result, localMin)
	}

	return result
}

// Sequential sliding window solution (optimal O(n))
func minSubArrayLenSequential(target int, nums []int) int {
	n := len(nums)
	minLen := math.MaxInt32
	left := 0
	sum := 0

	for right := 0; right < n; right++ {
		sum += nums[right]

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

type task struct {
	startIdx int
	endIdx   int
}

type result struct {
	minLen int
}

func minSubArrayLenVer4(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// For small arrays, use sequential
	if n < 100 {
		return minSubArrayLenSequentialVer2(target, nums)
	}

	numWorkers := 4
	chunkSize := max(n/numWorkers, 1)

	// Create task and result channels
	tasks := make(chan task, numWorkers)
	results := make(chan result, numWorkers)

	// Start worker pool
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(target, nums, tasks, results, &wg)
	}

	// Send tasks
	go func() {
		for i := 0; i < n; i += chunkSize {
			end := min(i+chunkSize, n)
			tasks <- task{startIdx: i, endIdx: end}
		}
		close(tasks)
	}()

	// Close results channel after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	minLen := math.MaxInt32
	for res := range results {
		minLen = min(minLen, res.minLen)
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func worker(target int, nums []int, tasks <-chan task, results chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range tasks {
		minLen := math.MaxInt32

		// Process chunk using sliding window
		for left := t.startIdx; left < t.endIdx; left++ {
			sum := 0
			for right := left; right < len(nums); right++ {
				sum += nums[right]
				if sum >= target {
					minLen = min(minLen, right-left+1)
					break
				}
			}
		}

		results <- result{minLen: minLen}
	}
}

func minSubArrayLenSequentialVer2(target int, nums []int) int {
	n := len(nums)
	minLen := math.MaxInt32
	left := 0
	sum := 0

	for right := 0; right < n; right++ {
		sum += nums[right]

		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
