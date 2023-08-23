package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0

	// Find the maximum number of candies among all kids
	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}

	// Check if each kid can have the maximum number of candies
	result := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}
