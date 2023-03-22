//70. Climbing Stairs
//Companies
//You are climbing a staircase. It takes n steps to reach the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//
//
//Example 1:
//
//Input: n = 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps
//Example 2:
//
//Input: n = 3
//Output: 3
//Explanation: There are three ways to climb to the top.
//1. 1 step + 1 step + 1 step
//2. 1 step + 2 steps
//3. 2 steps + 1 step
//
//
//Constraints:
//
//1 <= n <= 45

package main

import "log"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	result := make([]int, n+1)
	result[1] = 1
	result[2] = 2

	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
		log.Println("Solution", result[i], " - ", result[i-1], " - ", result[i-2])
	}

	return result[n]
}

func main() {
	log.Println(climbStairs(1))
	log.Println(climbStairs(2))
	log.Println(climbStairs(3))
	log.Println(climbStairs(4))
	log.Println(climbStairs(5))
	log.Println(climbStairs(6))

}
