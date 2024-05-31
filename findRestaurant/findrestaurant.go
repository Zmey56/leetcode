package findRestaurant

//Given two arrays of strings list1 and list2, find the common strings with the least index sum.
//
//A common string is a string that appeared in both list1 and list2.
//
//A common string with the least index sum is a common string such that if it appeared at list1[i] and list2[j]
//then i + j should be the minimum value among all the other common strings.
//
//Return all the common strings with the least index sum. Return the answer in any order.

func findRestaurant(list1 []string, list2 []string) []string {
	result := []string{}
	if len(list1) == 0 || len(list2) == 0 {
		return result
	}

	// Create a map to store the index of each string in list1
	indexMap := make(map[string]int)
	for i, v := range list1 {
		indexMap[v] = i
	}

	// Initialize the minimum index sum
	minIndexSum := len(list1) + len(list2)
	for j, v := range list2 {
		if i, ok := indexMap[v]; ok {
			if i+j < minIndexSum {
				minIndexSum = i + j
				result = []string{v}
			} else if i+j == minIndexSum {
				result = append(result, v)
			}
		}
	}

	return result
}
