package largestTriangleArea

import (
	"math"
	"runtime"
	"sync"
)

//Given an array of points on the X-Y plane points where points[i] = [xi, yi], return the area of the largest triangle
//that can be formed by any three different points. Answers within 10-5 of the actual answer will be accepted.

func largestTriangleArea(points [][]int) float64 {
	// The area of a triangle with vertices (x1, y1), (x2, y2), and (x3, y3) is given by:
	// 0.5 * abs(x1(y2 - y3) + x2(y3 - y1) + x3(y1 - y2))
	// We can iterate over all possible triangles and calculate the area of each one.
	// The time complexity of this approach is O(n^3), where n is the length of the points array.
	// However, we can improve this to O(n^2) by using the shoelace formula to calculate the area of a triangle.
	// The shoelace formula states that the area of a triangle with vertices (x1, y1), (x2, y2), and (x3, y3) is given by:
	// 0.5 * abs(x1(y2 - y3) + x2(y3 - y1) + x3(y1 - y2))
	// We can iterate over all possible pairs of points and calculate the area of the triangle formed by the pair and each other point.
	// We can then return the maximum area found.
	// The time complexity of this approach is O(n^2), where n is the length of the points array.
	// The space complexity is O(1).
	maxArea := 0.0
	for i, p1 := range points {
		for j, p2 := range points {
			if i == j {
				continue
			}
			for k, p3 := range points {
				if i == k || j == k {
					continue
				}
				area := 0.5 * abs(p1[0]*(p2[1]-p3[1])+p2[0]*(p3[1]-p1[1])+p3[0]*(p1[1]-p2[1]))
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea

}

func abs(x int) float64 {
	if x < 0 {
		return float64(-x)
	}
	return float64(x)
}

func largestTriangleAreaVer2(points [][]int) float64 {
	numGoroutines := runtime.NumCPU()
	areas := make([]float64, numGoroutines)

	var wg sync.WaitGroup
	partSize := len(points) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		if i == numGoroutines-1 {
			end = len(points)
		}

		wg.Add(1)
		go func(part [][]int, i int) {
			defer wg.Done()
			for j, p1 := range part {
				for k, p2 := range part {
					if j == k {
						continue
					}
					for l, p3 := range part {
						if j == l || k == l {
							continue
						}
						area := 0.5 * math.Abs(float64(p1[0]*(p2[1]-p3[1])+p2[0]*(p3[1]-p1[1])+p3[0]*(p1[1]-p2[1])))
						if area > areas[i] {
							areas[i] = area
						}
					}
				}
			}
		}(points[start:end], i)
	}

	wg.Wait()

	maxArea := 0.0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
