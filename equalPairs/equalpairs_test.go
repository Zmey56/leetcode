package main

import (
	"testing"
)

func TestCountEqualPairs(t *testing.T) {
	// Test case 1
	grid1 := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}
	expected1 := 1
	result1 := equalPairs(grid1)
	if result1 != expected1 {
		t.Errorf("Test case 1: Expected %d, but got %d", expected1, result1)
	}

	// Test case 2
	grid2 := [][]int{
		{3, 1, 2, 2},
		{1, 4, 4, 5},
		{2, 4, 2, 2},
		{2, 4, 2, 2},
	}
	expected2 := 3
	result2 := equalPairs(grid2)
	if result2 != expected2 {
		t.Errorf("Test case 2: Expected %d, but got %d", expected2, result2)
	}

	grid3 := [][]int{
		{11, 1},
		{1, 11},
	}
	expected3 := 2
	result3 := equalPairs(grid3)
	if result3 != expected3 {
		t.Errorf("Test case 2: Expected %d, but got %d", expected3, result3)
	}

}
