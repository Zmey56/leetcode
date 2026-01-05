package findMinArrowShots

import (
	"slices"
	"sort"
)

// 452. Minimum Number of Arrows to Burst Balloons

//There are some spherical balloons taped onto a flat wall that represents the XY-plane. The balloons are represented as
//a 2D integer array points where points[i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches between
//xstart and xend. You do not know the exact y-coordinates of the balloons.
//
//Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis.
//A balloon with xstart and xend is burst by an arrow shot at x if xstart <= x <= xend. There is no limit to the number of
//arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path.
//
//Given the array points, return the minimum number of arrows that must be shot to burst all balloons.

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	currentEnd := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] > currentEnd {
			arrows++
			currentEnd = points[i][1]
		}
	}

	return arrows

}

func findMinArrowShotsV2(points [][]int) int {
	n := len(points)
	if n <= 1 {
		return n
	}

	// Используем современный slices.SortFunc (быстрее чем sort.Slice)
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1] // Осторожно с переполнением int, но для LeetCode ок
	})

	arrows := 1
	currentEnd := points[0][1]

	for i := 1; i < n; i++ {
		// Используем прямые обращения к индексам, минимизируем аллокации
		if points[i][0] > currentEnd {
			arrows++
			currentEnd = points[i][1]
		}
	}

	return arrows
}
