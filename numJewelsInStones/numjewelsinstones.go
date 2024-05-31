package numJewelsInStones

import (
	"runtime"
	"sync"
)

// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have.
//Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

// Letters are case sensitive, so "a" is considered a different type of stone from "A".
func numJewelsInStones(jewels string, stones string) int {
	jewelsMap := make(map[rune]struct{})
	for _, j := range jewels {
		jewelsMap[j] = struct{}{}
	}

	count := 0
	for _, s := range stones {
		if _, ok := jewelsMap[s]; ok {
			count++
		}
	}

	return count
}

func numJewelsInStonesVer2(jewels string, stones string) int {
	jewelsMap := make(map[rune]struct{})
	for _, j := range jewels {
		jewelsMap[j] = struct{}{}
	}

	numGouroutines := runtime.NumCPU()
	counts := make(chan int, numGouroutines)

	var wg sync.WaitGroup
	partSize := len(stones) / numGouroutines

	for i := 0; i < numGouroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			count := 0
			start := i * partSize
			end := start + partSize
			if i == numGouroutines-1 {
				end = len(stones)
			}
			for _, s := range stones[start:end] {
				if _, ok := jewelsMap[s]; ok {
					count++
				}
			}
			counts <- count
		}(i)
	}

	go func() {
		wg.Wait()
		close(counts)
	}()

	total := 0
	for c := range counts {
		total += c
	}

	return total
}
