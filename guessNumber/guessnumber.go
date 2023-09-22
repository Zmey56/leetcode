package main

import "fmt"

func guess(num int) int {
	pickedNumber := 6 // Change this to the actual picked number
	if num < pickedNumber {
		return -1 // Your guess is lower
	} else if num > pickedNumber {
		return 1 // Your guess is higher
	} else {
		return 0 // Your guess is correct
	}
}

func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2
		result := guess(mid)

		if result == 0 {
			return mid // Found the picked number
		} else if result == -1 {
			right = mid - 1 // Adjust the right boundary
		} else {
			left = mid + 1 // Adjust the left boundary
		}
	}

	return -1 // Picked number not found (this should not happen)
}

func main() {
	n1, pick1 := 10, 6
	fmt.Println("Example 1:", guessNumber(n1), pick1) // Output: 6

	n2, pick2 := 1, 1
	fmt.Println("Example 2:", guessNumber(n2), pick2) // Output: 1

	n3, pick3 := 2, 1
	fmt.Println("Example 3:", guessNumber(n3), pick3) // Output: 1
}
