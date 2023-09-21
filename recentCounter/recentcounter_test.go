package main

import (
	"testing"
)

func TestRecentCounter(t *testing.T) {
	recentCounter := Constructor()

	// Test case 1
	if result := recentCounter.Ping(1); result != 1 {
		t.Errorf("Expected 1, but got %d", result)
	}

	// Test case 2
	if result := recentCounter.Ping(100); result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}

	// Test case 3
	if result := recentCounter.Ping(3001); result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}

	// Test case 4
	if result := recentCounter.Ping(3002); result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}
}
