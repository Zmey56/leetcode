package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	ages := []int{10, 15, 30, 7, 4, 34, 6}
	fmt.Println(sortListVer2(ages))
}

func sortList(ages []int) []int {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})

	return ages

}
func sortListVer2(ages []int) []int {
	chunkSize := len(ages) / 4

	var chucks [][]int
	for i := 0; i < len(ages); i += chunkSize {
		end := i + chunkSize
		if end > len(ages) {
			end = len(ages)
		}
		chucks = append(chucks, ages[i:end])
	}

	var wg sync.WaitGroup
	wg.Add(len(chucks))
	for _, chunk := range chucks {
		go func(c []int) {
			defer wg.Done()
			sort.Ints(c)
		}(chunk)
	}
	wg.Wait()

	sortedDat := merge(chucks)

	return sortedDat

}

func merge(chucks [][]int) []int {
	var result []int
	for _, chuck := range chucks {
		result = append(result, chuck...)
	}
	return result
}

func sortListZeroAndOne(val []int) []int {
	mapZeroOne := make(map[int]int)
	for _, v := range val {
		mapZeroOne[v]++
	}

	result := make([]int, len(val)-1)
	zeroSlice := make([]int, mapZeroOne[0])
	oneSlice := make([]int, mapZeroOne[1])
	for key, count := range mapZeroOne {
		if key == 0 {
			for i := 0; i < count; i++ {
				zeroSlice = append(zeroSlice, 0)
			}
		} else {
			for i := 0; i < count; i++ {
				oneSlice = append(zeroSlice, 1)
			}
		}
	}

	result = append(zeroSlice, oneSlice...)

	return result
}
