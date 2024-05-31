package findRestaurant

import "testing"

// Test cases for the findRestaurant function
func Test_findRestaurant(t *testing.T) {
	// Test case where list1 and list2 are empty
	list1 := []string{}
	list2 := []string{}
	expected := []string{}
	result := findRestaurant(list1, list2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where list1 is empty
	list1 = []string{}
	list2 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	expected = []string{}
	result = findRestaurant(list1, list2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where list2 is empty
	list1 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	list2 = []string{}
	expected = []string{}
	result = findRestaurant(list1, list2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where list1 and list2 have no common strings
	list1 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	list2 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	expected = []string{}
	result = findRestaurant(list1, list2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}

	// Test case where list1 and list2 have one common string
	list1 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}
	list2 = []string{"KFC", "Burger King", "Tapioca Express", "Shogun", "KFC"}
	expected = []string{"KFC"}
	result = findRestaurant(list1, list2)
	if !equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

// Helper function to check if two slices are equal
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
