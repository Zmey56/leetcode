package relativeSortArray

import (
	"sort"
)

//1122. Relative Sort Array

//Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.
//
//Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2.
//Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.

func relativeSortArray(arr1 []int, arr2 []int) []int {
	mapArr1 := make(map[int]int)
	for _, v := range arr1 {
		mapArr1[v]++
	}

	result := make([]int, 0)
	for _, v := range arr2 {
		for i := 0; i < mapArr1[v]; i++ {
			result = append(result, v)
		}
		delete(mapArr1, v)
	}

	rest := make([]int, 0)
	for k, v := range mapArr1 {
		for i := 0; i < v; i++ {
			rest = append(rest, k)
		}
	}

	sort.Ints(rest)

	result = append(result, rest...)

	return result
}

func relativeSortArrayVer2(arr1 []int, arr2 []int) []int {
	maxVal := 0
	for _, num := range arr1 {
		if num > maxVal {
			maxVal = num
		}
	}

	count := make([]int, maxVal+1)
	for _, num := range arr1 {
		count[num]++
	}

	result := make([]int, 0, len(arr1))
	for _, num := range arr2 {
		for count[num] > 0 {
			result = append(result, num)
			count[num]--
		}
	}

	for num, c := range count {
		for c > 0 {
			result = append(result, num)
			c--
		}
	}

	return result
}

func relativeSortArrayVer3(arr1 []int, arr2 []int) []int {
	orderMap := make(map[int]int)
	for index, value := range arr2 {
		orderMap[value] = index
	}

	// Channels for concurrently sorting elements
	ch := make(chan []int, 2)

	// Goroutine to sort elements present in arr2
	go func() {
		arr := make([]int, 0)
		for _, num := range arr1 {
			if _, exists := orderMap[num]; exists {
				arr = append(arr, num)
			}
		}
		sort.Slice(arr, func(i, j int) bool {
			return orderMap[arr[i]] < orderMap[arr[j]]
		})
		ch <- arr
	}()

	// Goroutine to sort elements not present in arr2
	go func() {
		arr := make([]int, 0)
		for _, num := range arr1 {
			if _, exists := orderMap[num]; !exists {
				arr = append(arr, num)
			}
		}
		sort.Ints(arr)
		ch <- arr
	}()

	// Combine the results from both goroutines
	part1 := <-ch
	part2 := <-ch

	return append(part2, part1...)
}
