package main

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		asteroid []int
		expected []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
	}

	for _, test := range tests {
		result := asteroidCollision(test.asteroid)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For asteroid %v, expected %v, but got %v", test.asteroid, test.expected, result)
		}
	}
}
