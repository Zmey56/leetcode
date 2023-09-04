package main

import "fmt"

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, asteroid := range asteroids {
		for asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			last := stack[len(stack)-1]

			if last == -asteroid {
				stack = stack[:len(stack)-1]
				asteroid = 0
			} else if last > -asteroid {
				asteroid = 0
			} else {
				stack = stack[:len(stack)-1]
			}
		}

		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

func main() {
	// Test cases
	asteroids1 := []int{5, 10, -5}
	fmt.Println(asteroidCollision(asteroids1)) // Output: [5 10]

	asteroids2 := []int{8, -8}
	fmt.Println(asteroidCollision(asteroids2)) // Output: []

	asteroids3 := []int{10, 2, -5}
	fmt.Println(asteroidCollision(asteroids3)) // Output: [10]

	asteroids4 := []int{8, -8, -8, 8}
	fmt.Println(asteroidCollision(asteroids4)) // Output: []
}
